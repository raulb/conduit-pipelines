#!/bin/bash

confluent-hub install --no-prompt debezium/debezium-connector-postgresql:latest

su -c "/etc/confluent/docker/run" appuser
