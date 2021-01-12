package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"lastfmsearch/cmd/lastfmsearch/config"
	"os"
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

// main Main
func main() {
	fmt.Println("It's alive!")
}
