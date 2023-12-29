package esgenerics

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
)

type QueryMap map[string]interface{}

// QueryFn EsQueryFn is a function that constructs an Elasticsearch query.
type QueryFn func() QueryMap

// Store defines the interface for Elasticsearch operations.
//
//go:generate mockery --name=Store --inpackage --case=snake
type Store[E Entity[M], M any] interface {
	ExecuteSearch(ctx context.Context, handler QueryHandler) (*Result[E, M], error)
}

// ElasticsearchStore manages Elasticsearch operations.
type ElasticsearchStore[E Entity[M], M any] struct {
	client    ElasticSearchClient
	indexName string
}

// NewStore creates a new instance of ElasticsearchStore.
func NewStore[E Entity[M], M any](client ElasticSearchClient, indexName string) *ElasticsearchStore[E, M] {
	return &ElasticsearchStore[E, M]{client: client, indexName: indexName}
}

// ExecuteSearch performs a search query in Elasticsearch and returns the results.
func (store *ElasticsearchStore[E, M]) ExecuteSearch(ctx context.Context, handler QueryHandler) (*Result[E, M], error) {
	query := handler.BuildQuery(ctx)
	resp, err := store.client.Search(ctx, store.indexName, query)
	if err != nil {
		return nil, fmt.Errorf("error executing search query: %w", err)
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	var searchResult *Result[E, M]
	if err = json.Unmarshal(responseBody, &searchResult); err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %w", err)
	}

	return searchResult, nil
}
