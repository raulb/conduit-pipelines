COMPOSE_FILE = compose.yml

.PHONY: setup
setup:
	docker-compose -f $(COMPOSE_FILE) up -d

.PHONY: clean
clean:
	docker-compose -f $(COMPOSE_FILE) down --volumes --remove-orphans
	rm -rf plugins

.PHONY: deploy-sink-connector
deploy-sink-connector:
	curl -X POST -H "Content-Type: application/json" --data @sink-snowflake.json http://localhost:8083/connectors

.PHONY: kafka-connect-logs
kafka-connect-logs:
	docker logs -f kafka-connect

.PHONY: produce-messages
produce-messages:
	../produce-messages.sh

.PHONY: open-control-center
open-control-center:
	open http://localhost:9021/clusters

.PHONY: check-localstack-connection
check-localstack-connection:
	curl -d '{}' snowflake.localhost.localstack.cloud:4566/session