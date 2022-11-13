package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/tomas-mota/go-bitbucket"
)

const (
	projectName = "TestProject"
	projectKey  = "TPO"
	desc        = "My Test Project"
	updatedDesc = "My Updated Project"
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

	// Create
	_, err = c.Projects.CreateProject(ctx,
		&bitbucket.CreateProjectRequest{
			Name:        projectName,
			Key:         projectKey,
			Description: desc,
			Public:      true,
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("project created")

	// Update
	p, err := c.Projects.UpdateProject(ctx,
		&bitbucket.UpdateProjectRequest{
			Key:         projectKey,
			Description: updatedDesc,
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	if p.Description != updatedDesc {
		log.Fatal("project was not updated as expected")
	}
	fmt.Println("project updated")

	// Get and validate again
	p, err = c.Projects.GetProject(ctx,
		&bitbucket.GetProjectRequest{
			Key: projectKey,
		},
	)

	if err != nil {
		log.Fatal(err)
	}

	if p.Description != updatedDesc {
		log.Fatal("project was not updated as expected")
	}

	fmt.Println("project update validated")

	//Delete
	err = c.Projects.DeleteProject(ctx,
		&bitbucket.DeleteProjectRequest{
			Key: projectKey,
		},
	)

	if err != nil {
		panic(err)
	}

	fmt.Println("project deleted")

}
