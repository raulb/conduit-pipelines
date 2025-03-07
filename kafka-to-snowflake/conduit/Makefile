include .env

COMPOSE_FILE = compose.yml
TOPIC=source-topic
BROKER=broker:29092

.PHONY: setup
setup:
	docker-compose -f $(COMPOSE_FILE) up -d source-kafka

.PHONY: start-conduit
start-conduit:
	docker-compose -f $(COMPOSE_FILE) up -d conduit-with-snowflake

.PHONY: clean
clean:
	docker-compose -f $(COMPOSE_FILE) down --volumes --remove-orphans

.PHONY: connectkafka
connectkafka:
	docker exec -it source-kafka /bin/sh

.PHONY: create-topic
create-topic:
	docker exec -it source-kafka kafka-topics --create --topic $(TOPIC) --bootstrap-server $(BROKER) --partitions 1 --replication-factor 1

.PHONY: produce
produce:
	docker exec -it source-kafka /usr/bin/kafka-console-producer --topic $(TOPIC) --bootstrap-server $(BROKER)

.PHONY: consume
consume:
	docker exec -it source-kafka /usr/bin/kafka-console-consumer --bootstrap-server $(BROKER) --topic $(TOPIC) --from-beginning

.PHONY: list-topics
list-topics:
	docker exec -it source-kafka /usr/bin/kafka-topics --list --bootstrap-server $(BROKER)

.PHONY: reset-topics
reset-topics:
	docker exec -it source-kafka /bin/sh -c "\
	  /usr/bin/kafka-topics --list --bootstrap-server $(BROKER) | \
	  xargs -I {} kafka-topics --delete --topic {} --bootstrap-server $(BROKER)"

.PHONY: conduit-logs
conduit-logs:
	docker logs -f conduit-with-snowflake

.PHONY: open-snowflake
open-snowflake:
	open -u https://$(SNOWFLAKE_HOST)

.PHONY: produce-messages
produce-messages:
	./produce-messages.sh 