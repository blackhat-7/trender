package sources

import (
	"github.com/google/go-github/v76/github"
)

type GitHubClient struct {
	client *github.Client
}

func NewGitHubClient() *GitHubClient {
	// Use installation transport with client.
	client := github.NewClient(nil).WithAuthToken("... your access token ...")
	return &GitHubClient{
		client: client,
	}
}
