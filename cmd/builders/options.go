package builders

import (
	"log"
	"time"
)

type (
	dbClient struct {
		url         string
		connections int
		timeout     time.Duration
	}

	option func(*dbClient)
)

func WithUrl(url string) option {
	if url == "" {
		log.Fatal("url cannot be empty")
	}

	return func(db *dbClient) {
		db.url = url
	}
}

func WithConnections(connections int) option {
	if connections < 1 {
		log.Fatal("connections must be greater than 1")
	}

	return func(db *dbClient) {
		db.connections = connections
	}
}

func WithTimeout(timeout time.Duration) option {
	return func(db *dbClient) {
		db.timeout = timeout
	}
}

func NewDBClient(opts ...option) *dbClient {
	db := &dbClient{
		timeout: 5 * time.Second,
	}

	for _, opt := range opts {
		opt(db)
	}

	return db
}

func (db *dbClient) Connect() {
	log.Printf("Connecting to %s with %d connections and timeout %s\n", db.url, db.connections, db.timeout)
	// simulate connection logic here
	time.Sleep(1 * time.Second)
	log.Println("Connected!")
}
