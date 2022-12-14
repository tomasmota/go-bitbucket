package bitbucket

import (
	"context"
	"fmt"
)

// ProjectService provides operations around bitbucket projects
type ProjectService interface {
	GetProject(context.Context, *GetProjectRequest) (*Project, error)
	CreateProject(context.Context, *CreateProjectRequest) (*Project, error)
	DeleteProject(context.Context, *DeleteProjectRequest) error
	UpdateProject(context.Context, *UpdateProjectRequest) (*Project, error)
	AddPermission(context.Context, *AddPermissionRequest) error
}

type projectService struct {
	client *Client
}

// Project represents a Bitbucket Project
type Project struct {
	Name        string `json:"name"`
	Key         string `json:"key"`
	ID          int    `json:"id"`
	Description string `json:"description"`
	Scope       string `json:"scope,omitempty"`
	Type        string `json:"type"`
	Public      bool   `json:"public"`
}

// GetProjectRequest contains the fields required to fetch a project
type GetProjectRequest struct {
	Key string `json:"key" validate:"required"`
}

func (ps *projectService) GetProject(ctx context.Context, getReq *GetProjectRequest) (*Project, error) {
	err := validate.Struct(getReq)
	if err != nil {
		return nil, ErrParameters
	}
	req, err := ps.client.newRequest("GET", fmt.Sprintf("projects/%s", getReq.Key), nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request for getting projects: %w", err)
	}

	p := Project{}
	err = ps.client.do(ctx, req, &p)
	if err != nil {
		return nil, fmt.Errorf("error fetching projects: %w", err)
	}
	return &p, nil
}

// CreateProjectRequest contains the fields required to create a project
type CreateProjectRequest struct {
	Name        string `json:"name" validate:"required"`
	Key         string `json:"key" validate:"required"`
	Description string `json:"description,omitempty"`
	Public      bool   `json:"public,omitempty"`
}

func (ps *projectService) CreateProject(ctx context.Context, createReq *CreateProjectRequest) (*Project, error) {
	err := validate.Struct(createReq)
	if err != nil {
		return nil, ErrParameters
	}

	req, err := ps.client.newRequest("POST", "projects", createReq)
	if err != nil {
		return nil, fmt.Errorf("error creating request for creating project: %w", err)
	}

	p := Project{}
	err = ps.client.do(ctx, req, &p)
	if err != nil {
		return nil, fmt.Errorf("error creating project: %w", err)
	}

	return &p, nil
}

// DeleteProjectRequest contains the fields required to delete a project
type DeleteProjectRequest struct {
	Key string `json:"key" validate:"required"`
}

func (ps *projectService) DeleteProject(ctx context.Context, deleteReq *DeleteProjectRequest) error {
	err := validate.Struct(deleteReq)
	if err != nil {
		return ErrParameters
	}
	req, err := ps.client.newRequest("DELETE", fmt.Sprintf("projects/%s", deleteReq.Key), nil)
	if err != nil {
		return fmt.Errorf("error creating request for deleting project: %w", err)
	}

	err = ps.client.do(ctx, req, nil)
	if err != nil {
		return fmt.Errorf("error deleting project: %w", err)
	}

	return nil
}

// UpdateProjectRequest contains the fields required to update a project
type UpdateProjectRequest struct {
	Key         string `json:"key" validate:"required"`
	Description string `json:"description,omitempty"`
	Public      bool   `json:"public,omitempty"`
}

func (ps *projectService) UpdateProject(ctx context.Context, updateReq *UpdateProjectRequest) (*Project, error) {
	err := validate.Struct(updateReq)
	if err != nil {
		return nil, ErrParameters
	}
	req, err := ps.client.newRequest("PUT", fmt.Sprintf("projects/%s", updateReq.Key), updateReq)
	if err != nil {
		return nil, fmt.Errorf("error creating request for updating project: %w", err)
	}

	p := Project{}
	err = ps.client.do(ctx, req, &p)
	if err != nil {
		return nil, fmt.Errorf("error updating project: %w", err)
	}

	return &p, nil
}

// AddPermissionRequest contains the fields required to give a certain group permissions in a project
type AddPermissionRequest struct {
	ProjectKey string `validate:"required"`
	Group      string `validate:"required"`
	Permission string `validate:"required,oneof='PROJECT_READ' 'PROJECT_WRITE' 'PROJECT_ADMIN'"`
}

func (ps *projectService) AddPermission(ctx context.Context, apReq *AddPermissionRequest) error {
	err := validate.Struct(apReq)
	if err != nil {
		return ErrParameters
	}

	req, err := ps.client.newRequest("PUT", fmt.Sprintf("projects/%s/permissions/groups", apReq.ProjectKey), apReq)
	if err != nil {
		return fmt.Errorf("error creating request for adding permission to project: %w", err)
	}
	q := req.URL.Query()
	q.Add("name", apReq.Group)
	q.Add("permission", apReq.Permission)

	err = ps.client.do(ctx, req, nil)
	if err != nil {
		return fmt.Errorf("error adding permission to project: %w", err)
	}

	return nil
}

// RevokePermissionRequest contains the fields required to revoke a certain group permissions in a project
type RevokePermissionRequest struct {
	ProjectKey string `validate:"required"`
	Group      string `validate:"required"`
}

func (ps *projectService) RevokePermission(ctx context.Context, apReq *RevokePermissionRequest) error {
	err := validate.Struct(apReq)
	if err != nil {
		return ErrParameters
	}

	req, err := ps.client.newRequest("PUT", fmt.Sprintf("projects/%s/permissions/groups", apReq.ProjectKey), apReq)
	if err != nil {
		return fmt.Errorf("error creating request for adding permission to project: %w", err)
	}
	q := req.URL.Query()
	q.Add("name", apReq.Group)

	err = ps.client.do(ctx, req, nil)
	if err != nil {
		return fmt.Errorf("error adding permission to project: %w", err)
	}

	return nil
}
