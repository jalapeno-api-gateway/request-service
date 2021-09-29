package redis

import (
	"context"

	"github.com/jalapeno-api-gateway/model/class"
)

func Fetch(ctx context.Context, keys []string, className class.Class) []interface{} {
	keys = prependCollectionNameToKeys(keys, className)
	var documents []interface{}
	values := getValuesByKeys(ctx, keys)
	for _, value := range values {
		documents = append(documents, unmarshalObject(value, className))
	}
	return documents
}

func FetchAll(ctx context.Context, className class.Class) []interface{} {
	var documents []interface{}
	keys := scanAllKeysOfCollection(ctx, className)
	values := getValuesByKeys(ctx, keys)
	for _, value := range values {
		documents = append(documents, unmarshalObject(value, className))
	}
	return documents
}