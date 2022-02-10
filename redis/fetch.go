package redis

import (
	"context"

	"github.com/jalapeno-api-gateway/jagw-core/model/class"
	"github.com/sirupsen/logrus"
)

func Fetch(ctx context.Context, logger *logrus.Entry, keys []string, className class.Class) []interface{} {
	keys = prependCollectionNameToKeys(keys, className)
	var documents []interface{}
	values := getValuesByKeys(ctx, logger, keys)
	for _, value := range values {
		documents = append(documents, unmarshalObject(logger, value, className))
	}
	return documents
}

func FetchAll(ctx context.Context, logger *logrus.Entry, className class.Class) []interface{} {
	var documents []interface{}
	keys := scanAllKeysOfCollection(ctx, className)
	values := getValuesByKeys(ctx, logger, keys)
	for _, value := range values {
		documents = append(documents, unmarshalObject(logger, value, className))
	}
	return documents
}