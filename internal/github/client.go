package github

import (
	"fmt"

	githubv3 "github.com/google/go-github/v86/github"
)

type Client struct {
	Raw   *githubv3.Client
	Token string
	Owner string
	Repo  string
}

func NewClient(token, owner, repo, ghesBaseURL string) (*Client, error) {
	if token == "" {
		return nil, fmt.Errorf("github: token is required")
	}
	var raw *githubv3.Client
	if ghesBaseURL != "" {
		c, err := githubv3.NewEnterpriseClient(ghesBaseURL, "", nil)
		if err != nil {
			return nil, fmt.Errorf("github: new enterprise client: %w", err)
		}
		raw = c
	} else {
		raw = githubv3.NewClient(nil)
	}
	return &Client{Raw: raw.WithAuthToken(token), Token: token, Owner: owner, Repo: repo}, nil
}
