package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/jalapeno-api-gateway/jagw-core/model/class"
	"github.com/jalapeno-api-gateway/jagw-core/model/topology"
)

func prependCollectionNameToKeys(keys []string, className class.Class) []string {
	for i, key := range keys {
		keys[i] = fmt.Sprintf("%s/%s", className, key)
	}
	return keys
}

func scanAllKeysOfCollection(ctx context.Context, className class.Class) []string {
	iter := RedisClient.Scan(ctx, 0, fmt.Sprintf("%s/*", className), 0).Iterator()
	keys := []string{}
	for iter.Next(ctx) {
		keys = append(keys, iter.Val())
	}
	return keys
}

func getValuesByKeys(ctx context.Context, keys []string) [][]byte {
	//MGet returns nil for a key which is not present in cache
	values, err := RedisClient.MGet(ctx, keys...).Result()
	if err != nil {
		log.Fatal("Error fetching documents from Redis: ", err)
	}

	bytes := [][]byte{}
	for _, value := range values {
		if value != nil { //entry found in cache
			bytes = append(bytes, []byte(value.(string)))
		}
	}

	return bytes
}

func unmarshalObject(bytes []byte, className class.Class) interface{} {
	switch className {
		case class.LSNode:
			document := topology.LSNode{}
			handleUnmarshallingError(json.Unmarshal(bytes, &document))
			return document
		case class.LSLink:
			document := topology.LSLink{}
			handleUnmarshallingError(json.Unmarshal(bytes, &document))
			return document
		case class.LSPrefix:
			document := topology.LSPrefix{}
			handleUnmarshallingError(json.Unmarshal(bytes, &document))
			return document
		case class.LSSRv6SID:
			document := topology.LSSRv6SID{}
			handleUnmarshallingError(json.Unmarshal(bytes, &document))
			return document
		default: return nil
	}
}

func handleUnmarshallingError(err error) {
	if err != nil {
		log.Fatal("Error while unmarshalling object: ", err)
	}
}