## Conduit Pipelines

This repository will contain some Conduit Pipelines that will be used to easily test its performance.

> [!NOTE]  
> Until > 0.13.2 is released (see PR), you'll need to build the Conduit image locally
> `$ (conduitio/conduit) docker buildx build --platform linux/arm64 -t ghcr.io/conduitio/conduit:v0.13.1 --load ~/code/conduitio/conduit`

### Postgres -> Kafka


1. Start PostgresQL and Kafka `docker-compose -f ./pg-to-kafka/docker-compose.yml up -d pg-0 broker control-center`
1. Insert data `make insert-data`
1. Start Conduit `docker buildx build --platform linux/arm64 -t ghcr.io/conduitio/conduit:v0.13.1 --load ~/code/conduitio/conduit && docker-compose -f ./pg-to-kafka/docker-compose.yml up -d conduit`
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

##### Clean up

```bash
make clean && rm -Rf pg-to-kafka/data/
```

### Kafka -> Snowflake

TODO


