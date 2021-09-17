package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
)

func prependCollectionNameToKeys(keys []string, collectionName CollectionName) []string {
	for i, key := range keys {
		keys[i] = fmt.Sprintf("%s/%s", collectionName, key)
	}
	return keys
}

func scanAllKeysOfCollection(ctx context.Context, collectionName CollectionName) []string {
	iter := RedisClient.Scan(ctx, 0, fmt.Sprintf("%s/*", collectionName), 0).Iterator()
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

//
// ---> Unmarshalling <---
//

func unmarshalLsNodeDocument(bytes []byte) LsNodeDocument {
	document := LsNodeDocument{}
	err := json.Unmarshal(bytes, &document)
	if err != nil {
		log.Fatal("Error unmarshalling LsNode: ", err)
	}
	return document
}

func unmarshalLsLinkDocument(bytes []byte) LsLinkDocument {
	document := LsLinkDocument{}
	err := json.Unmarshal(bytes, &document)
	if err != nil {
		log.Fatal("Error unmarshalling LsLink: ", err)
	}
	return document
}
