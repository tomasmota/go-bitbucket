package integration

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/tomas-mota/go-bitbucket"
)

func TestRepos_CRUD(t *testing.T) {
	const (
		projectName     = "TestProject"
		projectKey      = "TPO"
		repoName        = "TestRepo"
		updatedRepoName = "UpdatedTestRepo"
	)

	ctx := context.Background()

	// Create Project to put repo in
	_, err := client.Projects.CreateProject(ctx,
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

	// Create Repo
	r, err := client.Repos.CreateRepo(ctx,
		&bitbucket.CreateRepoRequest{
			ProjectKey: projectKey,
			Name:       repoName,
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	slug := r.Slug
	fmt.Printf("repo %s created\n", slug)

	// Update Repo
	_, err = client.Repos.UpdateRepo(ctx, projectKey, slug, &bitbucket.UpdateRepoRequest{Name: updatedRepoName})
	if err != nil {
		log.Fatal(err)
	}

	// Get Repo
	r, err = client.Repos.GetRepo(ctx, projectKey, slug)
	if err != nil {
		log.Fatal(err)
	}
	if r.Name != updatedRepoName {
		log.Fatal("repo was not updated")
	}
	fmt.Printf("repo %s has been correctly updated \n", slug)

	// Delete Repo
	err = client.Repos.DeleteRepo(ctx, projectKey, slug)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("repo %s deleted\n", slug)

	// Delete Project
	err = client.Projects.DeleteProject(ctx,
		&bitbucket.DeleteProjectRequest{
			Key: projectKey,
		},
	)

	if err != nil {
		panic(err)
	}

	fmt.Println("project deleted")
}
