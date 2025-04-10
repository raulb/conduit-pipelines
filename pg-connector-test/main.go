package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/conduitio/conduit-commons/config"
	"github.com/conduitio/conduit-connector-postgres"
	sdk "github.com/conduitio/conduit-connector-sdk"
)

func main() {
	ctx := context.Background()
	src := postgres.Connector.NewSource()
	cfg := config.Config{
		"tables":                             "employees",
		"url":                                "postgresql://meroxauser:meroxapass@localhost:5432/meroxadb",
		"cdcMode":                            "logrepl",
		"logrepl.slotName":                   "conduit_slot",
		"logrepl.publicationName":            "conduit_pub",
		"logrepl.autoCleanup":                "true",
		"logrepl.withAvroSchema":             "false",
		"snapshotMode":                       "initial", // Changed from "never" to "initial" to capture existing data
		"sdk.batch.size":                     "10000",
		"sdk.batch.delay":                    "0s",
		"sdk.schema.extract.key.enabled":     "false",
		"sdk.schema.extract.payload.enabled": "false",
	}
	err := sdk.Util.ParseConfig(ctx, cfg, src.Config(), postgres.Connector.NewSpecification().SourceParams)
	if err != nil {
		panic(fmt.Errorf("error parsing config: %v", err))
	}

	// Open the source with a nil position (start from beginning)
	err = src.Open(ctx, nil)
	if err != nil {
		panic(fmt.Errorf("error opening source: %v", err))
	}
	defer src.Teardown(ctx)

	// read first record (?)
	rec, err := src.Read(ctx)
	if errors.Is(err, sdk.ErrBackoffRetry) {
		fmt.Println("record not available")
	} else if err != nil {
		panic(fmt.Errorf("error reading from source: %v", err))
	} else {
		fmt.Println(rec)
	}

	// Execute the script to insert records
	// The script should take a command line argument for the number of records to insert
	recordsToInsert := "100000" // Default to 1000 records
	if len(os.Args) > 1 {
		recordsToInsert = os.Args[1]
	}

	recordsToInsertInt, nil := strconv.Atoi(recordsToInsert)
	log.Printf("Running script to insert %s records...", recordsToInsert)

	scriptPath := "./scripts/insert_named_employees.sh"

	// Start writing
	insertDuration, insertRate, err := writeRecords(scriptPath, recordsToInsert)
	if err != nil {
		log.Fatalf("Error executing script: %v", err)
	}
	log.Printf("Insertion completed in %v (%.2f records/second)", insertDuration, insertRate)

	// Measure performance of reading records
	log.Println("Starting to read records...")

	start := time.Now()

	for i := 0; i < recordsToInsertInt; i++ {
		next, err := src.Read(ctx)
		if err != nil {
			panic(fmt.Errorf("error reading from source: %v", err))
		}

		if i%10000 == 0 {
			elapsed := time.Since(start).Seconds()
			fmt.Printf("total count: %v, elapsed: %v, rate: %v\n", i, elapsed, float64((i))/elapsed)
			err := os.WriteFile("position.json", next.Position, 0644)
			if err != nil {
				panic(fmt.Errorf("failed to write position.json: %w", err))
			}
		}
	}

	duration := time.Now().Sub(start).Seconds()
	fmt.Println("duration:", duration)

	// Calculate and log performance metrics
	recordsPerSecond := float64(recordsToInsertInt) / duration

	log.Printf("Performance Summary:")
	log.Printf("- Total records read: %d", recordsToInsertInt)
	log.Printf("- Total read duration: %v", duration)
	log.Printf("- Read rate: %.2f records/second", recordsPerSecond)
}

func writeRecords(scriptPath, recordsToInsert string) (time.Duration, float64, error) {
	// Start timing the insert operation
	insertStartTime := time.Now()

	cmd := exec.Command("/bin/bash", scriptPath, recordsToInsert)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return 0, 0, fmt.Errorf("error executing script: %v\nOutput: %s", err, string(output))
	}

	insertDuration := time.Since(insertStartTime)
	insertRate := float64(0)
	// Try to calculate insertion rate if we can convert the record count
	if insertCount, err := strconv.Atoi(recordsToInsert); err == nil {
		insertRate = float64(insertCount) / insertDuration.Seconds()
	}

	log.Printf("Script output:\n%s", string(output))
	return insertDuration, insertRate, nil
}
