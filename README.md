# Conduit Pipelines

This repository will contain some Conduit Pipelines that will be used to easily test its performance.

## Postgres -> Kafka

### Conduit

> [!NOTE]  
> All commands would be executed from `./pg-to-kafka/conduit`.

1. Start PostgresQL and Kafka `docker-compose -f ./pg-to-kafka/conduit/compose.yml up -d pg-0 broker control-center`
1. Insert data `make insert-data`
1. Start Conduit `docker-compose -f ./pg-to-kafka/conduit/compose.yml up -d conduit`
1. Access control center


```bash
open http://localhost:9021/clusters
```

**Test data**

- 5 batches of 80k records each => 400000 records

```
meroxadb=# SELECT pg_size_pretty(pg_total_relation_size('employees')) AS size;
 size  
-------
 29 MB
(1 row)
```

**Results**

- Conduit: (CDC) 1,108MB/s 

### Kafka Connect


#### Clean up

```bash
make clean && rm -Rf pg-to-kafka/data/
```

## Kafka -> Snowflake

> [!NOTE]  
> All commands would be executed from `./kafka-to-snowflake`.
>
> In order to execute these tests, a [Snowflake account is required](https://signup.snowflake.com/) and provide those credentials in .env

1. `cp .env.sample .env`
1. Update credentials
1. `make setup`
1. To monitor execution on Conduit: `make conduit-logs`
1. Create topic (this is optional as topic will be created automatically): `make create-topic`
1. Produce messages: `make produce-messages`
1. Access Snowflake instance.
