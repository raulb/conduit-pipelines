COMPOSE_FILE = compose.yml

.PHONY: setup
setup:
	docker-compose -f $(COMPOSE_FILE) up -d

.PHONY: clean
clean:
	docker-compose -f $(COMPOSE_FILE) down --volumes --remove-orphans

.PHONY: deploy-source-connector
deploy-source-connector:
	curl -X POST -H "Content-Type: application/json" --data @source-postgres.json http://localhost:8083/connectors

.PHONY: kafka-connect-logs
kafka-connect-logs:
	docker logs -f kafka-connect

.PHONY: connectpg
connectpg:
	docker exec -it source-postgres psql -U meroxauser -d meroxadb

.PHONY: insert-data
insert-data:
	../insert-data.sh 

.PHONY: open-control-center
open-control-center:
	open http://localhost:9021/clusters