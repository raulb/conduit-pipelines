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

// Helper function to read records from the source connector
func readRecords(ctx context.Context, src sdk.Source) (int, error) {
	recordCount := 0
	for {
		record, err := src.Read(ctx)
		if err != nil {
			if errors.Is(sdk.ErrBackoffRetry, err) {
				log.Println("Backing off, waiting for more records...")
				time.Sleep(1 * time.Second)
				continue
			}
			// Consider any other error as end of records to keep things simple
			log.Printf("Stopped reading records due to: %v", err)
			if recordCount > 0 {
				// If we've read some records, just return them and treat this as completion
				return recordCount, nil
			}
			return recordCount, fmt.Errorf("error reading record: %w", err)
		}

		recordCount++
		log.Printf("Read record %d: %v", recordCount, record)
	}
}

func main() {
	ctx := context.Background()
	src := postgres.Connector.NewSource()

	cfg := config.Config{
		"tables":                             "employees",
		"url":                                "postgresql://meroxauser:meroxapass@test-pg-connector:5432/meroxadb",
		"cdcMode":                            "logrepl",
		"logrepl.slotName":                   "conduit_slot",
		"logrepl.publicationName":            "conduit_pub",
		"logrepl.autoCleanup":                "true",
		"logrepl.withAvroSchema":             "false",
		"snapshotMode":                       "never",
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

	// Execute the script to insert records
	// The script should take a command line argument for the number of records to insert
	recordsToInsert := "1000" // Default to 1000 records
	if len(os.Args) > 1 {
		recordsToInsert = os.Args[1]
	}
	log.Printf("Running script to insert %s records...", recordsToInsert)

	// Path to the script should be adjusted based on your environment
	scriptPath := "./scripts/insert_named_employees.sh"

	// Start timing the insert operation
	insertStartTime := time.Now()

	cmd := exec.Command("/bin/bash", scriptPath, recordsToInsert)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Error executing script: %v\nOutput: %s", err, string(output))
	}

	insertDuration := time.Since(insertStartTime)
	insertRate := float64(0)
	// Try to calculate insertion rate if we can convert the record count
	if insertCount, err := strconv.Atoi(recordsToInsert); err == nil {
		insertRate = float64(insertCount) / insertDuration.Seconds()
	}

	log.Printf("Script output:\n%s", string(output))
	log.Printf("Insertion completed in %v (%.2f records/second)",
		insertDuration, insertRate)

	// Measure performance of reading records
	log.Println("Starting to read records...")
	startTime := time.Now()

	recordCount, err := readRecords(ctx, src)
	if err != nil {
		log.Fatalf("Error reading records: %v", err)
	}

	duration := time.Since(startTime)

	// Calculate and log performance metrics
	recordsPerSecond := float64(recordCount) / duration.Seconds()

	log.Printf("Performance Summary:")
	log.Printf("- Total records read: %d", recordCount)
	log.Printf("- Total read duration: %v", duration)
	log.Printf("- Read rate: %.2f records/second", recordsPerSecond)

	// Try to convert to expected number for validation
	expectedRecords, convErr := strconv.Atoi(recordsToInsert)
	if convErr == nil {
		if recordCount != expectedRecords {
			log.Printf("Warning: Expected to read %d records, but actually read %d",
				expectedRecords, recordCount)
		} else {
			log.Printf("Success: Read all %d expected records", expectedRecords)
		}

		// Overall performance metrics
		log.Printf("Overall Performance:")
		log.Printf("- Insert rate: %.2f records/second", float64(expectedRecords)/insertDuration.Seconds())
		log.Printf("- Read rate: %.2f records/second", recordsPerSecond)
		totalDuration := insertDuration + duration
		log.Printf("- Total processing time: %v", totalDuration)
	}
}
