package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"lastfmsearch/cmd/lastfmsearch/config"
	"os"
	"log"
	"net/http"
	"lastfmsearch/pkg/graph"
	"lastfmsearch/pkg/graph/generated"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

var cfg config.Config

// init Initialize config from args or env variables
func init() {
	parser := flags.NewParser(&cfg, flags.Default)
	_, err := parser.Parse()
	if err != nil {
		if e, ok := err.(*flags.Error); !ok || e.Type != flags.ErrHelp {
			fmt.Printf("Initialization error: %s.\n", err)
		}
		os.Exit(1)
	}
}

// main Main func
func main() {
	fmt.Println("Starting lastfmsearch service...")

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	http.Handle("/query", srv)
	if cfg.EnablePlayground {
		http.Handle("/playground", playground.Handler("GraphQL playground", "/query"))
		log.Printf("connect to http://localhost:%s/playground for GraphQL playground", cfg.HttpPort)
	}
	log.Fatal(http.ListenAndServe(":"+cfg.HttpPort, nil))
}
