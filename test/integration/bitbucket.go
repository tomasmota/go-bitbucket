package integration

import (
	"log"
	"os"

	"github.com/tomas-mota/go-bitbucket"
)

var client *bitbucket.Client

// Configures a client to be used in testing
// NOTE: these integration tests require a bitbucket instance to be running on localhost
func init() {
	username := os.Getenv("BITBUCKET_USERNAME")
	if username == "" {
		log.Fatal("no username provide")
	}
	password := os.Getenv("BITBUCKET_PASSWORD")
	if password == "" {
		log.Fatal("no password provided")
	}

	c, err := bitbucket.NewClient(
		bitbucket.Config{
			Host:   "localhost",
			Scheme: "http",

			Username: username,
			Password: password,
		},
	)
	if err != nil {
		panic(err)
	}
	client = c
}
