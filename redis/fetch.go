package redis

import (
	"context"

	"github.com/jalapeno-api-gateway/jagw-core/jagwerror"
	"github.com/jalapeno-api-gateway/jagw-core/model/class"
	"github.com/sirupsen/logrus"
)

func Fetch(ctx context.Context, logger *logrus.Entry, keys []string, className class.Class) ([]interface{}, *jagwerror.Error) {
	keys = prependCollectionNameToKeys(keys, className)
	var documents []interface{}
	values, keysNotFoundError := getValuesByKeys(ctx, logger, keys)
	for _, value := range values {
		documents = append(documents, unmarshalObject(logger, value, className))
	}
	return documents, keysNotFoundError
}

func FetchAll(ctx context.Context, logger *logrus.Entry, className class.Class) ([]interface{}, *jagwerror.Error) {
	var documents []interface{}
	keys := scanAllKeysOfCollection(ctx, className)
	values, keysNotFoundError := getValuesByKeys(ctx, logger, keys)
	for _, value := range values {
		documents = append(documents, unmarshalObject(logger, value, className))
	}
	return documents, keysNotFoundError
}