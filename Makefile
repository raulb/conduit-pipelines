COMPOSE_FILE = containers/pg-to-kafka/docker-compose.yml

.PHONY: setup
setup:
	docker-compose -f $(COMPOSE_FILE) up -d

.PHONY: clean
clean:
	docker-compose -f $(COMPOSE_FILE) down --volumes --remove-orphans


.PHONY: create-table
create-table:
	docker exec -i pg-0 psql -U meroxauser -d meroxadb -f /docker-entrypoint-initdb.d/init.sql

.PHONY: insert-data
insert-data:
	./containers/pg-to-kafka/insert-data.sh 


# Help command to list available targets
.PHONY: help
help:
	@echo "Available targets:"
	@echo "  setup          - Setup Docker images and start containers"
	@echo "  clean          - Tear down Docker containers and remove images"
	@echo "  insert-data    - Insert data into Postgres"
	@echo "  help           - Display this help message"
