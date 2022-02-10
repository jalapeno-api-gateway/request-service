package requestservice

import (
	"context"

	"github.com/jalapeno-api-gateway/jagw-core/model/class"
	"github.com/jalapeno-api-gateway/request-service/redis"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/peer"
)

func fetchDocuments(ctx context.Context, logger *logrus.Entry, keys []string, className class.Class) []interface{} {
	var documents []interface{}
	if len(keys) == 0 {
		documents = redis.FetchAll(ctx, logger, className)
	} else {
		documents = redis.Fetch(ctx, logger, keys, className)
	}
	return documents
}

func getClientIp(ctx context.Context) string {
	p, status := peer.FromContext(ctx);
	if status {
		return p.Addr.String()
	}
	return ""
}