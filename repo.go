package bitbucket

import (
	"context"
	"fmt"
)

// RepoService provides operations around bitbucket repos
type RepoService interface {
	GetRepo(context.Context, *GetRepoRequest) (*Repo, error)
	CreateRepo(context.Context, *CreateRepoRequest) (*Repo, error)
	DeleteRepo(context.Context, *DeleteRepoRequest) error
	UpdateRepo(context.Context, *UpdateRepoRequest) (*Repo, error)
}

type repoService struct {
	client *Client
}

// TODO: set omitempty on optional fields
// Repo represents a Bitbucket Repo
type Repo struct {
	Project       *Project `json:"project"`
	Slug          string   `json:"slug"`
	Name          string   `json:"name"`
	Description   string   `json:"description"`
	HierarcyId    string   `json:"hierarchyId"`
	StatusMessage string   `json:"statusMessage"`
	Archived      bool     `json:"archive"`
	Forkable      bool     `json:"forkable"`
	DefaultBranch string   `json:"defaultBranch"`
	ScmId         string   `json:"scmId"`
	Scope         bool     `json:"scope"`
	Id            int      `json:"id"`
	State         string   `json:"state"`
	Public        bool     `json:"public"`
}

// GetRepoRequest contains the fields required to fetch a repo
type GetRepoRequest struct {
	ProjectKey string
	Slug       string
}

func (rs *repoService) GetRepo(ctx context.Context, getReq *GetRepoRequest) (*Repo, error) {
	req, err := rs.client.newRequest("GET", fmt.Sprintf("projects/%s/repos/%s", getReq.ProjectKey, getReq.Slug), nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request for getting repos: %w", err)
	}

	r := Repo{}
	err = rs.client.do(ctx, req, &r)
	if err != nil {
		return nil, fmt.Errorf("error fetching repos: %w", err)
	}
	return &r, nil
}

// TODO: Add other fields
// CreateRepoRequest contains the fields required to create a repo
type CreateRepoRequest struct {
	ProjectKey  string
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description,omitempty"`
}

func (rs *repoService) CreateRepo(ctx context.Context, createReq *CreateRepoRequest) (*Repo, error) {
	req, err := rs.client.newRequest("POST", fmt.Sprintf("projects/%s/repos", createReq.ProjectKey), createReq)
	if err != nil {
		return nil, fmt.Errorf("error creating request for creating repo: %w", err)
	}

	r := Repo{}
	err = rs.client.do(ctx, req, &r)
	if err != nil {
		return nil, fmt.Errorf("error creating repo: %w", err)
	}

	return &r, nil
}

// DeleteRepoRequest contains the fields required to delete a repo
type DeleteRepoRequest struct {
	ProjectKey string
	Slug       string `json:"slug"`
}

func (rs *repoService) DeleteRepo(ctx context.Context, deleteReq *DeleteRepoRequest) error {
	req, err := rs.client.newRequest("DELETE", fmt.Sprintf("projects/%s/repos/%s", deleteReq.ProjectKey, deleteReq.Slug), nil)
	if err != nil {
		return fmt.Errorf("error creating request for deleting repo: %w", err)
	}

	err = rs.client.do(ctx, req, nil)
	if err != nil {
		return fmt.Errorf("error deleting repo: %w", err)
	}

	return nil
}

// UpdateRepoRequest contains the fields required to update a repo
type UpdateRepoRequest struct {
	ProjectKey  string
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description,omitempty"`
}

func (rs *repoService) UpdateRepo(ctx context.Context, updateReq *UpdateRepoRequest) (*Repo, error) {
	req, err := rs.client.newRequest("PUT", fmt.Sprintf("projects/%s/repos/%s", updateReq.ProjectKey, updateReq.Slug), updateReq)
	if err != nil {
		return nil, fmt.Errorf("error creating request for updating repo: %w", err)
	}

	r := Repo{}
	err = rs.client.do(ctx, req, &r)
	if err != nil {
		return nil, fmt.Errorf("error updating repo: %w", err)
	}

	return &r, nil
}
