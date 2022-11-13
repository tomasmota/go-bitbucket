package main

import (
	"context"
	"log"
	"os"

	"github.com/tomas-mota/go-bitbucket"
)

func main() {
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

	ctx := context.Background()
	c.Projects.CreateProject(
		ctx,
		&bitbucket.CreateProjectRequest{
			Name:        "TestProject",
			Key:         "TPO",
			Description: "My Test Project",
			Public:      true,
		},
	)

	c.Projects.UpdateProject(
		ctx,
		&bitbucket.UpdateProjectRequest{
			Key:         "TPO",
			Description: "Updated Project",
		},
	)

	c.Projects.DeleteProject(
		ctx,
		&bitbucket.DeleteProjectRequest{
			Key: "TPO",
		},
	)

}
