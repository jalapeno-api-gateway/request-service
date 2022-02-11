module github.com/jalapeno-api-gateway/request-service

go 1.16

require (
	github.com/go-redis/redis/v8 v8.11.2
	github.com/iancoleman/strcase v0.2.0
	github.com/influxdata/influxdb1-client v0.0.0-20200827194710-b269163b24ab
	github.com/jalapeno-api-gateway/jagw-core v1.2.1-0.20220210203330-054df1dad886
	github.com/jalapeno-api-gateway/protorepo-jagw-go v1.2.0
	github.com/sirupsen/logrus v1.8.1
	golang.org/x/net v0.0.0-20210805182204-aaa1db679c0d // indirect
	golang.org/x/sys v0.0.0-20210809222454-d867a43fc93e // indirect
	google.golang.org/grpc v1.42.0
	google.golang.org/protobuf v1.27.1
)
