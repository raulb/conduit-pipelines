# Makefile for setting up Docker images and running Conduit pipelines

# Variables
COMPOSE_FILE = containers/docker-compose.yml
PIPELINE_FILE = pg-to-kafka/pipeline.yml

# Default target
.PHONY: all
all: setup

# Setup Docker images and start containers
.PHONY: setup
setup:
	docker-compose -f $(COMPOSE_FILE) up -d

# Tear down Docker containers and remove images
.PHONY: clean
clean:
	docker-compose -f $(COMPOSE_FILE) down --volumes --remove-orphans

# Initialize Conduit pipeline
.PHONY: init-pipeline
init-pipeline:
	docker exec conduit conduit pipelines apply /app/pipelines/pipeline.yml

# Insert data into Postgres
.PHONY: insert-data
insert-data:
	docker exec -i pg-0 psql -U meroxauser -d meroxadb -f /docker-entrypoint-initdb.d/init.sql

# Help command to list available targets
.PHONY: help
help:
	@echo "Available targets:"
	@echo "  setup          - Setup Docker images and start containers"
	@echo "  clean          - Tear down Docker containers and remove images"
	@echo "  init-pipeline  - Initialize Conduit pipeline"
	@echo "  insert-data    - Insert data into Postgres"
	@echo "  help           - Display this help message"
