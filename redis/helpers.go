package redis

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/jalapeno-api-gateway/jagw-core/jagwerror"
	"github.com/jalapeno-api-gateway/jagw-core/model/class"
	"github.com/jalapeno-api-gateway/jagw-core/model/topology"
	"github.com/sirupsen/logrus"
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

func getValuesByKeys(ctx context.Context, logger *logrus.Entry, keys []string) ([][]byte, *jagwerror.Error) {
	//MGet returns nil for a key which is not present in cache
	values, err := RedisClient.MGet(ctx, keys...).Result()
	if err != nil {
		logger.WithError(err).Panic("Failed to fetch documents from Redis.")
	}

	keysNotFound := []string{}
	bytes := [][]byte{}
	for i, value := range values {
		if value != nil { //entry found in cache
			bytes = append(bytes, []byte(value.(string)))
		} else {
			keysNotFound = append(keysNotFound, keys[i])
		}
	}
	
	return bytes, jagwerror.CreateErrorForKeysNotFound(keysNotFound)
}

func unmarshalObject(logger *logrus.Entry, bytes []byte, className class.Class) interface{} {
	switch className {
		case class.LsNode:
			document := topology.LsNode{}
			handleUnmarshallingError(logger, json.Unmarshal(bytes, &document))
			return document
		case class.LsNodeCoordinates:
			document := topology.LsNodeCoordinates{}
			handleUnmarshallingError(logger, json.Unmarshal(bytes, &document))
			return document
		case class.LsLink:
			document := topology.LsLink{}
			handleUnmarshallingError(logger, json.Unmarshal(bytes, &document))
			return document
		case class.LsPrefix:
			document := topology.LsPrefix{}
			handleUnmarshallingError(logger, json.Unmarshal(bytes, &document))
			return document
		case class.LsSrv6Sid:
			document := topology.LsSrv6Sid{}
			handleUnmarshallingError(logger, json.Unmarshal(bytes, &document))
			return document
		case class.LsNodeEdge:
			document := topology.LsNodeEdge{}
			handleUnmarshallingError(logger, json.Unmarshal(bytes, &document))
			return document
		default: return nil
	}
}

func handleUnmarshallingError(logger *logrus.Entry, err error) {
	if err != nil {
		logger.WithError(err).Panic("Failed to unmarshall object.")
	}
}