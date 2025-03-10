#!/bin/bash

cleanup() {
  echo "Interrupted. Exiting..."
  exit
}

trap cleanup SIGINT

batch_size=80000
batches=5

topic="source-topic" 

# Produce a batch of records similar to this.
# NOTE: is one of the records that the Kafka destination would receive from Postgres.
# {
#   "full_time": true,
#   "id": 399951,
#   "name": "John Doe",
#   "updated_at": "2025-03-04T11:58:13.910712Z"
# }

produce_batch() {
  for ((j=1; j<=batch_size; j++)); do
    id=$(( (i - 1) * batch_size + j ))
    updated_at=$(date -u +"%Y-%m-%dT%H:%M:%SZ")
    echo "{\"full_time\": true, \"id\": $id, \"name\": \"John Doe\", \"updated_at\": \"$updated_at\"}"
  done | docker exec -i source-kafka /usr/bin/kafka-console-producer --broker-list broker:29092 --topic "$topic"
}

start_time=$(date +%s)

for ((i=1; i<=batches; i++)); do
  echo "Producing batch $i..."
  produce_batch
done

end_time=$(date +%s)

duration=$((end_time - start_time))

echo "Produced $((batch_size * batches)) messages in $duration seconds."
