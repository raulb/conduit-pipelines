## Conduit Pipelines

This repository will contain some Conduit Pipelines that will be used to easily test its performance.

> [!NOTE]  
> Until > 0.13.2 is released (see PR), you'll need to build the Conduit image locally
> `$ (conduitio/conduit) docker buildx build --platform linux/arm64 -t ghcr.io/conduitio/conduit:v0.13.1 --load .`

### Postgres -> Kafka


1. Setup containers

```bash
make setup
```

2. Create schema

```bash
make insert-data
```

If this doesn't work, ensure schema was successfully created:

```bash
docker exec -it pg-0 psql -U meroxauser -d meroxadb
psql (16.4 (Debian 16.4-1.pgdg110+2))
Type "help" for help.

meroxadb=# \d
                 List of relations
 Schema |       Name       |   Type   |   Owner    
--------+------------------+----------+------------
 public | employees        | table    | meroxauser
 public | employees_id_seq | sequence | meroxauser

 ```
 
 3. Access control center

```bash
open http://localhost:9021/clusters
```

### Kafka -> Snowflake

TODO


