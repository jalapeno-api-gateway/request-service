package redis

import (
	"context"
)

//
// ---> FETCH SPECIFIC <---
//

func FetchLsNodes(ctx context.Context, keys []string) []LsNodeDocument {
	keys = prependCollectionNameToKeys(keys, LsNodeCollection)
	documents := []LsNodeDocument{}
	values := getValuesByKeys(ctx, keys)
	for _, value := range values {
		documents = append(documents, unmarshalLsNodeDocument(value))
	}
	return documents
}

func FetchLsLinks(ctx context.Context, keys []string) []LsLinkDocument {
	keys = prependCollectionNameToKeys(keys, LsLinkCollection)
	documents := []LsLinkDocument{}
	values := getValuesByKeys(ctx, keys)
	for _, value := range values {
		documents = append(documents, unmarshalLsLinkDocument(value))
	}
	return documents
}

//
// ---> FETCH ALL <---
//

func FetchAllLsNodes(ctx context.Context) []LsNodeDocument {
	documents := []LsNodeDocument{}
	keys := scanAllKeysOfCollection(ctx, LsNodeCollection)
	values := getValuesByKeys(ctx, keys)
	for _, value := range values {
		documents = append(documents, unmarshalLsNodeDocument(value))
	}
	return documents
}

func FetchAllLsLinks(ctx context.Context) []LsLinkDocument {
	documents := []LsLinkDocument{}
	keys := scanAllKeysOfCollection(ctx, LsLinkCollection)
	values := getValuesByKeys(ctx, keys)
	for _, value := range values {
		documents = append(documents, unmarshalLsLinkDocument(value))
	}
	return documents
}