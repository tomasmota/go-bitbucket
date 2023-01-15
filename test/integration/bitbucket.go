package integration

import (
	"github.com/tomas-mota/go-bitbucket"
)

var client *bitbucket.Client

// Configures a client to be used in testing
// NOTE: these integration tests require a bitbucket instance to be running on localhost
// Username and password should be "admin"
func init() {
	c, err := bitbucket.NewClient(
		bitbucket.Config{

			Host:   "localhost:7990",
			Scheme: "http",

			Username: "admin",
			Password: "admin",
		},
	)
	if err != nil {
		panic(err)
	}
	client = c
}
