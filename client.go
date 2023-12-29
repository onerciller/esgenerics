package esgenerics

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"os"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
)

// Options is used to configure the elasticsearch client
type Options struct {
	Username string
	Password string
	Hosts    []string
}

// ElasticSearchClient is the interface for elasticsearch client
//
//go:generate mockery --name=ElasticSearchClient --inpackage --case=snake
type ElasticSearchClient interface {
	Search(ctx context.Context, indexName string, queryMap QueryMap) (*esapi.Response, error)
}

type EsClient struct {
	client *elasticsearch.Client
	logger *log.Logger
}

func NewESClient(options ...func(*Options)) (*EsClient, error) {
	o := &Options{}

	for _, opt := range options {
		opt(o)
	}

	cfg := elasticsearch.Config{
		Addresses: o.Hosts,
		Username:  o.Username,
		Password:  o.Password,
	}

	client, err := elasticsearch.NewClient(cfg)

	if err != nil {
		return nil, err
	}

	logger := log.New(os.Stdout, "esgenerics: ", log.LstdFlags)
	return &EsClient{client: client, logger: logger}, nil
}

// Search executes a search query on the given index
func (es *EsClient) Search(ctx context.Context, indexName string, queryMap QueryMap) (*esapi.Response, error) {
	startTime := time.Now()

	es.logger.Printf("Debug: Executing search on index %s: %v", indexName, queryMap)
	var b bytes.Buffer
	err := json.NewEncoder(&b).Encode(queryMap)

	if err != nil {
		es.logger.Printf("Error executing search: %v", err)
		return nil, err
	}
	resp, err := es.client.Search(
		es.client.Search.WithContext(ctx),
		es.client.Search.WithIndex(indexName),
		es.client.Search.WithBody(&b),
	)
	es.logger.Printf("Debug: Search executed in %v", time.Since(startTime))
	return resp, err
}

// WithUsername sets the username for the elasticsearch client
func WithUsername(username string) func(*Options) {
	return func(o *Options) {
		o.Username = username
	}
}

// WithPassword sets the password for the elasticsearch client
func WithPassword(password string) func(*Options) {
	return func(o *Options) {
		o.Password = password
	}
}

// WithHosts sets the hosts for the elasticsearch client
func WithHosts(hosts string) func(*Options) {
	return func(o *Options) {
		o.Hosts = strings.Split(hosts, ",")
	}
}
