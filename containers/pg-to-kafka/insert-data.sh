#!/bin/bash

cleanup() {
  echo "Interrupted. Exiting..."
  exit
}

trap cleanup SIGINT

end=$((SECONDS+60))

# repeat the insert query for 60 seconds
while [ $SECONDS -lt $end ]; do
  docker exec -i pg-0 psql -U meroxauser -d meroxadb -c "
  INSERT INTO employees (name, full_time, updated_at) VALUES
  ('John Doe', true, NOW());
  "
done