package builders

import (
	"log"
	"time"
)

type dbClientBuilder struct {
	dbClient *dbClient
}

func (builder *dbClientBuilder) WithURL(url string) *dbClientBuilder {
	if url == "" {
		log.Fatalf("url cannot be empty")
	}
	builder.dbClient.url = url
	return builder
}

func (builder *dbClientBuilder) WithConnections(connections int) *dbClientBuilder {
	if connections < 1 {
		log.Fatalf("connections must be greater than 1")
	}
	builder.dbClient.connections = connections
	return builder
}

func (builder *dbClientBuilder) WithTimeout(timeout time.Duration) *dbClientBuilder {
	builder.dbClient.timeout = timeout
	return builder
}

func (builder *dbClientBuilder) Build() *dbClient {
	return builder.dbClient
}

func NewDBClientFluent() *dbClientBuilder {
	return &dbClientBuilder{
		dbClient: &dbClient{timeout: 5 * time.Second},
	}
}
