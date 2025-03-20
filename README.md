# Conduit Pipelines

This repository will contain some Conduit Pipelines that will be used to easily test its performance.

## Postgres -> Kafka

### Using Benchi

Having installed [`benchi`](https://github.com/conduitio/benchi) >= 0.3 is required.

```bash
go install github.com/conduitio/benchi/cmd/benchi@latest
```

#### CDC

```
make benchi-pg-to-kafka-cdc
```

### Conduit

> [!NOTE]  
> All commands would be executed from `./pg-to-kafka/conduit`.

#### Snapshot

1. Start everything but Conduit: `make setup`
1. Insert data `make insert-data`
1. Start Conduit `make start-conduit`
1. Access control center: `make open-control-center`

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

#### CDC

1. `make setup`
1. Deploy connector: `make deploy-source-connector`
1. Insert data: `make insert-data`
1. Access control center: `make open-control-center`

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

1. `make create-private-key`
1. `make show-private-key`
1. Copy that private key and create a user on your Snowflake cloud account with that one. Instructions [here](https://docs.confluent.io/cloud/current/connectors/cc-snowflake-sink/cc-snowflake-sink.html#creating-a-user-and-adding-the-public-key).
1. Make sure you grant a role to that user that has access to the database and table you're going to test against.
1. `make setup`
1. `cp sink-snowflake.sample.json sink-snowflake.json` and update all the proper values.
1. Make sure kafka connect is running and deploy your sink connector `make deploy-sink-connector`.
1. Produce messages `make produce-messages`.

Once finished, you should be able to see a new table on your configured database with the name `topic-ID`.
