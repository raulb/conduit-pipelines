#!/bin/bash

apt-get update && apt-get install -y jq curl wget

confluent-hub install --no-prompt debezium/debezium-connector-postgresql:latest

su -c "/etc/confluent/docker/run" appuser
