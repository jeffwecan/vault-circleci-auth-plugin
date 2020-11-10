package main

import (
	"net/url"

	circleci "github.com/tylux/go-circleci"
)

// CCIClient is an implementation of the Client interface that actually sends
// API calls to CircleCI.
type CCIClient struct {
	c *circleci.Client

	vcsType string
	owner   string
}

// NewCCIClient creates a Client instance with the provided token and returns it.
func NewCCIClient(token, vcsType, owner string) *CCIClient {
	return &CCIClient{
		c: &circleci.Client{
			Token: token,
			Debug: true,
		},
		vcsType: vcsType,
		owner:   owner,
	}
}

// GetBuild ...
func (c *CCIClient) GetBuild(project string, buildNum int) (*circleci.Build, error) {
	return c.c.GetBuild(c.vcsType, c.owner, project, buildNum)
}

// SetBaseURL ...
func (c *CCIClient) SetBaseURL(baseURL *url.URL) {
	c.c.BaseURL = baseURL
}
