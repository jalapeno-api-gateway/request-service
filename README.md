# request-service
The request-service is part of the Jalapeño API Gateway. It takes simple requests from SR-Apps and fetches the data from the feeders.

## gRPC
- When the file `proto/request-service/request-service.proto` is updated, this command needs to be run to recompile the code:
```bash
$ /bin/protoc/bin/protoc --proto_path=./proto/request-service --go_out=./proto/request-service --go_opt=paths=source_relative --go-grpc_out=./proto/request-service --go-grpc_opt=paths=source_relative ./proto/request-service/request-service.proto
```
- When the file `proto/graph-db-feeder/graph-db-feeder.proto` is updated, this command needs to be run to recompile the code:
```bash
$ /bin/protoc/bin/protoc --proto_path=./proto/graph-db-feeder --go_out=./proto/graph-db-feeder --go_opt=paths=source_relative --go-grpc_out=./proto/graph-db-feeder --go-grpc_opt=paths=source_relative ./proto/graph-db-feeder/graph-db-feeder.proto
```

## Setting Up Development Environment
Make sure you have setup the [global development environment](https://gitlab.ost.ch/ins/jalapeno-api/request-service/-/wikis/Development-Environment) first.

### Step 1: Initialize Okteto
- Clone the repository:
```bash
$ git clone ssh://git@gitlab.ost.ch:45022/ins/jalapeno-api/request-service.git
```
- Initialize okteto:
```bash
$ okteto init
```
- Replace content of okteto.yml with the following:
```yml
name: request-service
autocreate: true
image: okteto/golang:1
command: bash
namespace: jagw-dev-<namespace-name>
securityContext:
  capabilities:
    add:
      - SYS_PTRACE
volumes:
  - /go/pkg/
  - /root/.cache/go-build/
  - /root/.vscode-server
  - /go/bin/
  - /bin/protoc/
sync:
  - .:/usr/src/app
forward:
  - 2347:2345
  - 8082:8080
environment:
  - GRAPH_DB_FEEDER_ADDRESS=graph-db-feeder:9001
  - TSDB_FEEDER_ADDRESS=tsdb-service:9000

```

### Step 2: Initialize the Container
- Open VSCode in the root of the repository.
- Hit `cmd`  + `p` to open the command pallet.
- Enter `>` and then choose `okteto up`
- When prompted, choose your `okteto.yml` file.
- When prompted, choose `Linux` as the containers operating system.

### Step 3: Setup the Container
- In the VSCode instance from the container, install the `Go` extension, otherwise the command `go` will not work on the VSCode command line.
- Install any additional extensions you want.

#### Install the Protocol Buffer Compiler
Here is the official guide: https://grpc.io/docs/protoc-installation/  
Just run these commands:
```bash
$ apt update
$ apt install unzip
$ wget https://github.com/protocolbuffers/protobuf/releases/download/v3.17.3/protoc-3.17.3-linux-x86_64.zip
$ unzip protoc-3.17.3-linux-x86_64.zip -d /bin/protoc
$ rm protoc-3.17.3-linux-x86_64.zip
```

#### Install the gRPC Library for Go
Here is the official guide: https://grpc.io/docs/languages/go/quickstart/  
Just run these commands:
```bash
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
```



# Documentation
## Deployment Diagram
<img src="https://gitlab.ost.ch/ins/jalapeno-api/request-service/-/raw/master/assets/api-gateway-deployment.png" alt="deployment-diagram" width="800"/>

## Basic Principle
### For the Topology Data
- The data processor 'Topology' updates the GraphDB when required and upon completion sends an '_event' message to Kafka containing the '_key' of the updated database entry.
- During bootup, the GraphDB Feeder first fetches the current state of all collections in the GraphDB and stores everything in the cache.
- Afterwards, the GraphDB Feeder turns to Kafka and watches all '_event' topics. Whenever a database entry has been updated, the GraphDB Feeder reads the updated entry from the GraphDB (using the '_key') and compares it to what it has stored in the cache. It then sends and update to the Push Service containing only the updated properties and the '_key'. Finally, the GraphDB Feeder updates the cache.

This guarantees that the cache is always up to date and that the Push Service gets informed on all updates.

### For the Time Series Data
- Because there is a huge amount of Time Series Data and because it is constantly chaning, it will not be cached.
- To get the current state of a TSDB entry, the Request Service or the Push Service call the TSDB Feeder, which queries the database for them.
- The TSDB Feeder constantly watches the Kafka topic 'jalapeno.telemetry' and informs the Push Service on all updates.

### For the Request Service
- If the request service receives a request for topology data, it reads it directly from the cache.
- If the request service receives a request for telemetry data, it asks the TSDB Feeder to get it.

### For the Push Service
- When an SR-App subscribes to some data, the Push Service first opens a channels to catch all incoming updates (either from the GraphDB Feeder or the TSDB Feeder).
- Then it queries the Request Service for the current state of these entries and returns them to the SR-App.
- Finally, it starts processing the updates in the channel, sending them all back to the SR-App. Some of the first updates may be redundant, but it is guaranteed that no update has been lost to the SR-App.


## Startup Behavior
This is the behavior of the API Gateway when it is started for the first time:
1. The GraphDB Feeder starts queueing all updates from the '_events' topics in Kafka concerning the topology data in a channel.
2. It then gets the current state of all entries from  the GraphDB and stores them in the cache.
3. It now starts processing the updates in the channel. Some of the first updates may be redundant, but it is guaranteed that no update has been lost.

