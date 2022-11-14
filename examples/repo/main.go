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
	repoName    = "TestRepo"
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

	// Create Project to put repo in
	_, err = c.Projects.CreateProject(ctx,
		&bitbucket.CreateProjectRequest{
			Name:   projectName,
			Key:    projectKey,
			Public: true,
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("project created")

	r, err := c.Repos.CreateRepo(ctx,
		&bitbucket.CreateRepoRequest{
			ProjectKey: projectKey,
			Name:       repoName,
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	slug := r.Slug
	fmt.Printf("repo %s created\n", r.Slug)

	// Delete Repo
	err = c.Repos.DeleteRepo(ctx,
		&bitbucket.DeleteRepoRequest{
			ProjectKey: projectKey,
			Slug:       slug,
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("repo %s deleted\n", r.Slug)

	// Delete Project
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
