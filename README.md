# request-service
The request-service is part of the Jalapeño API Gateway. It takes simple requests from SR-Apps and fetches the data from the feeders.

## gRPC
- When the file `proto/request-service/requestservice.proto` is updated, this command needs to be run to recompile the code:
```bash
$ protoc --proto_path=./proto/requestservice --go_out=./proto/requestservice --go_opt=paths=source_relative --go-grpc_out=./proto/requestservice --go-grpc_opt=paths=source_relative ./proto/requestservice/requestservice.proto
```

## Setting Up Development Environment
Make sure you have setup the [global development environment](https://github.com/jalapeno-api-gateway/request-service/-/wikis/Development-Environment) first.

## Initialize Okteto
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
namespace: <namespace-name>
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
  - APP_SERVER_ADDRESS=0.0.0.0:9000
  - INFLUX_DB_URL=http://10.20.1.24:30308
  - INFLUX_USER=root
  - INFLUX_PASSWORD=jalapeno
  - INFLUX_DB=mdt_db
  - REDIS_PASSWORD=a-very-complex-password-here
  - SENTINEL_ADDRESS=sentinel.<namespace-name>.svc.cluster.local:5000
  - SENTINEL_MASTER=mymaster
```
