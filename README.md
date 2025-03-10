# Conduit Pipelines

This repository will contain some Conduit Pipelines that will be used to easily test its performance.

## Postgres -> Kafka

**Test data**

- 5 batches of 80k records each => 400000 records

```
meroxadb=# SELECT pg_size_pretty(pg_total_relation_size('employees')) AS size;
 size  
-------
 29 MB
(1 row)
```

### Conduit

> [!NOTE]  
> All commands would be executed from `./pg-to-kafka/conduit`.

#### Snapshot

1. Start everything but Conduit: `make setup`
1. Insert data `make insert-data`
1. Start Conduit `make start-conduit`
1. Access control center: `make open-control-center`

**Results:**

#### CDC

1. Start everything but Conduit `make setup`
1. Start Conduit `make start-conduit`
1. Insert data `make insert-data`
1. Access control center: `make open-control-center`

### Kafka Connect

> [!NOTE]  
> All commands would be executed from `./pg-to-kafka/kafka-connect`.


#### Snapshot

1. `make setup`
1. Insert data: `make insert-data`
1. Deploy connector: `make deploy-source-connector`
1. Access control center: `make open-control-center`

**Results:**

#### CDC

1. `make setup`
1. Deploy connector: `make deploy-source-connector`
1. Insert data: `make insert-data`
1. Access control center: `make open-control-center`

**Results:**

## Kafka -> Snowflake

> [!NOTE]  
>
> In order to execute these tests, a [Snowflake account is required](https://signup.snowflake.com/) and provide those credentials in .env (`cp ./kafka-to-snowflake/.env.sample ./kafka-to-snowflake/.env`)
 

### Conduit

> [!NOTE]  
> All commands would be executed from `./kafka-to-snowflake/conduit`.

For every test, data would be inserted to the same table so bear in mind when checking the destination.

#### Snapshot

1. Setup everything but Conduit: `make setup`
1. Create topic (this is optional as topic will be created automatically): `make create-topic`
1. Produce messages: `make produce-messages`
1. Start Conduit: `make start-conduit` (`make conduit-logs` to monitor its logs)
1. Access Snowflake instance.

#### CDC

1. Setup everything but Conduit: `make setup`
1. Start Conduit: `make start-conduit` (`make conduit-logs` to monitor its logs)
1. Create topic (this is optional as topic will be created automatically): `make create-topic`
1. Produce messages: `make produce-messages`
1. Access Snowflake instance.

### Kafka Connect

> [!NOTE]  
> All commands would be executed from `./kafka-to-snowflake/kafka-connect`.

WIP
