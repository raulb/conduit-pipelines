COMPOSE_FILE = docker-compose.yml
TOPIC=source-topic

.PHONY: setup
setup:
	docker-compose -f $(COMPOSE_FILE) up -d

.PHONY: clean
clean:
	docker-compose -f $(COMPOSE_FILE) down --volumes --remove-orphans

.PHONY: connectkafka
connectkafka:
	docker exec -it source-kafka /bin/sh

.PHONY: create-topic
create-topic:
	docker exec -it source-kafka kafka-topics --create --topic $(TOPIC) --bootstrap-server broker:29092 --partitions 1 --replication-factor 1

.PHONY: produce
produce:
	docker exec -it source-kafka kafka-console-producer --topic $(TOPIC) --bootstrap-server broker:29092

.PHONY: consume
consume:
	docker exec -it source-kafka kafka-console-consumer --topic $(TOPIC) --bootstrap-server broker:29092 --from-beginning

.PHONY: reset-topics
reset-topics:
	docker exec -it source-kafka /bin/sh -c "\
	  kafka-topics --list --bootstrap-server broker:29092 | \
	  xargs -I {} kafka-topics --delete --topic {} --bootstrap-server broker:29092"