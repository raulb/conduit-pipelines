COMPOSE_FILE = pg-to-kafka/docker-compose.yml
CONDUIT_PATH=~/code/conduitio/conduit

.PHONY: setup
setup:
	docker buildx build --platform linux/arm64 -t ghcr.io/conduitio/conduit:v0.13.1 --load $(CONDUIT_PATH)
	docker-compose -f $(COMPOSE_FILE) up -d

.PHONY: clean
clean:
	docker-compose -f $(COMPOSE_FILE) down --volumes --remove-orphans

.PHONY: connectpg
connectpg:
	docker exec -it pg-0 psql -U meroxauser -d meroxadb

.PHONY: connectkafka
connectkafka:
	docker exec -it broker /bin/sh

.PHONY: insert-data
insert-data:
	./pg-to-kafka/insert-data.sh 
