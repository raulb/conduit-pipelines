COMPOSE_FILE = compose.yml

.PHONY: setup
setup:
	docker-compose -f $(COMPOSE_FILE) up -d

.PHONY: clean
clean:
	docker-compose -f $(COMPOSE_FILE) down --volumes --remove-orphans

.PHONY: connectpg
connectpg:
	docker exec -it source-postgres psql -U meroxauser -d meroxadb

.PHONY: connectkafka
connectkafka:
	docker exec -it destination-kafka /bin/sh

.PHONY: insert-data
insert-data:
	../insert-data.sh 
