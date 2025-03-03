COMPOSE_FILE = docker-compose.yml

.PHONY: setup
setup:
	docker-compose -f $(COMPOSE_FILE) up -d

.PHONY: clean
clean:
	docker-compose -f $(COMPOSE_FILE) down --volumes --remove-orphans

.PHONY: connectkafka
connectkafka:
	docker exec -it broker /bin/sh

.PHONY: create-topic
create-topic:
	docker exec -it broker kafka-topics --create --topic my-topic --bootstrap-server broker:29092 --partitions 1 --replication-factor 1

.PHONY: produce
produce:
	docker exec -it broker kafka-console-producer --topic my-topic --bootstrap-server broker:29092

.PHONY: consume
consume:
	docker exec -it broker kafka-console-consumer --topic my-topic --bootstrap-server broker:29092 --from-beginning

.PHONY: reset-topics
reset-topics:
	docker exec -it broker /bin/sh -c "\
	  kafka-topics --list --bootstrap-server broker:29092 | \
	  xargs -I {} kafka-topics --delete --topic {} --bootstrap-server broker:29092"