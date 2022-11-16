package integration

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/tomas-mota/go-bitbucket"
)

func TestProjects_CRUD(t *testing.T) {
	const (
		projectName = "TestProject"
		projectKey  = "TPO"
		desc        = "My Test Project"
		updatedDesc = "My Updated Project"
	)

	ctx := context.Background()

	// Create
	_, err := client.Projects.CreateProject(ctx,
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
	p, err := client.Projects.UpdateProject(ctx,
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
	p, err = client.Projects.GetProject(ctx,
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
