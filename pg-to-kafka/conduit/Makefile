COMPOSE_FILE = compose.yml

.PHONY: setup
setup:
	docker-compose -f $(COMPOSE_FILE) up -d source-postgres destination-kafka destination-control-center

.PHONY: start-conduit
start-conduit:
	docker-compose -f $(COMPOSE_FILE) up -d conduit

.PHONY: clean
clean:
	docker-compose -f $(COMPOSE_FILE) down --volumes --remove-orphans
	rm -rf data

.PHONY: connectpg
connectpg:
	docker exec -it source-postgres psql -U meroxauser -d meroxadb

.PHONY: connectkafka
connectkafka:
	docker exec -it destination-kafka /bin/sh

.PHONY: insert-data
insert-data:
	../insert-data.sh 

.PHONY: conduit-logs
conduit-logs:
	docker logs -f conduit

.PHONY: open-control-center
open-control-center:
	open http://localhost:9021/clusters