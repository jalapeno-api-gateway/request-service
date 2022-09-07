module github.com/jalapeno-api-gateway/request-service

go 1.16

require (
	github.com/go-redis/redis/v8 v8.11.2
	github.com/influxdata/influxdb1-client v0.0.0-20200827194710-b269163b24ab
	github.com/jalapeno-api-gateway/jagw-core v1.3.3-0.20220317160815-cba70f15248b
	github.com/jalapeno-api-gateway/protorepo-jagw-go v1.3.3-0.20220907094222-c6b0ca7e20b7
	github.com/nqd/flat v0.1.1
	github.com/sirupsen/logrus v1.8.1
	golang.org/x/net v0.0.0-20210805182204-aaa1db679c0d // indirect
	golang.org/x/sys v0.0.0-20210809222454-d867a43fc93e // indirect
	google.golang.org/grpc v1.49.0
	google.golang.org/protobuf v1.28.1
)
