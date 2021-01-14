package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/jessevdk/go-flags"
	"github.com/xlab/closer"
	"lastfmsearch/cmd/lastfmsearch/config"
	"lastfmsearch/pkg/graph"
	"lastfmsearch/pkg/graph/generated"
	"lastfmsearch/pkg/lastfm"
	"log"
	"net/http"
	"os"
)

var cfg config.Config

// init Initialize config from args or env variables
func init() {
	parser := flags.NewParser(&cfg, flags.Default)
	_, err := parser.Parse()
	if err != nil {
		if e, ok := err.(*flags.Error); !ok || e.Type != flags.ErrHelp {
			log.Printf("Initialization error: %s.\n", err)
		}
		os.Exit(1)
	}
}

// main Main func
func main() {
	log.Println("Starting lastfmsearch service...")
	closer.Bind(cleanup)

	client := lastfm.NewClient(cfg.LastfmApiEndpoint, cfg.LastfmApiKey, 1)
	resolver := graph.NewResolver(client)
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))
	http.Handle("/query", srv)
	if cfg.EnablePlayground {
		http.Handle("/playground", playground.Handler("GraphQL playground", "/query"))
		log.Printf("connect to http://localhost:%s/playground for GraphQL playground", cfg.HttpPort)
	}
	log.Fatal(http.ListenAndServe(":"+cfg.HttpPort, nil))
	closer.Hold()
}

// cleanup Callback on panic or system signal received
func cleanup() {
	log.Println("Service was stopped.")
}
