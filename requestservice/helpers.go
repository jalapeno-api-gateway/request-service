package requestservice

import (
	"context"

	"github.com/jalapeno-api-gateway/jagw-core/jagwerror"
	"github.com/jalapeno-api-gateway/jagw-core/model/class"
	"github.com/jalapeno-api-gateway/request-service/redis"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/peer"
)

func fetchDocuments(ctx context.Context, logger *logrus.Entry, keys []string, className class.Class) ([]interface{}, *jagwerror.Error) {
	var documents []interface{}
	var keysNotFoundError *jagwerror.Error
	if len(keys) == 0 {
		logger.Trace("No keys provided, fetching all documents.")
		documents, keysNotFoundError  = redis.FetchAll(ctx, logger, className)
	} else {
		documents, keysNotFoundError = redis.Fetch(ctx, logger, keys, className)
	}
	return documents, keysNotFoundError
}

func getClientIp(ctx context.Context) string {
	p, status := peer.FromContext(ctx);
	if status {
		return p.Addr.String()
	}
	return ""
}