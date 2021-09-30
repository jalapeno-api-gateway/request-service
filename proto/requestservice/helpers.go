package requestservice

import (
	"context"

	"github.com/jalapeno-api-gateway/jagw-core/model/class"
	"github.com/jalapeno-api-gateway/request-service/redis"
)

func fetchDocuments(ctx context.Context, keys []string, className class.Class) []interface{} {
	var documents []interface{}
	if len(keys) == 0 {
		documents = redis.FetchAll(ctx, className)
	} else {
		documents = redis.Fetch(ctx, keys, className)
	}
	return documents
}