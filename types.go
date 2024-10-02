package main

import (
	"context"

	"github.com/google/go-github/v65/github"
	"github.com/stretchr/testify/mock"
)

type GitHubClient interface {
	ListByOrg(ctx context.Context, org string, opt *github.RepositoryListByOrgOptions) ([]*github.Repository, *github.Response, error)
	ListAllTopics(ctx context.Context, owner string, repo string) ([]string, *github.Response, error)
	ReplaceAllTopics(ctx context.Context, owner string, repo string, topics []string) ([]string, *github.Response, error)
}

type clientWrapper struct {
	client *github.Client
}

type MockGitHubClient struct {
	mock.Mock
}

func (c clientWrapper) ListByOrg(ctx context.Context, org string, opt *github.RepositoryListByOrgOptions) ([]*github.Repository, *github.Response, error) {
	return c.client.Repositories.ListByOrg(ctx, org, opt)
}

func (c clientWrapper) ListAllTopics(ctx context.Context, owner string, repo string) ([]string, *github.Response, error) {
	return c.client.Repositories.ListAllTopics(ctx, owner, repo)
}

func (c clientWrapper) ReplaceAllTopics(ctx context.Context, owner string, repo string, topics []string) ([]string, *github.Response, error) {
	return c.client.Repositories.ReplaceAllTopics(ctx, owner, repo, topics)
}

func (m *MockGitHubClient) ListByOrg(ctx context.Context, org string, opt *github.RepositoryListByOrgOptions) ([]*github.Repository, *github.Response, error) {
	args := m.Called(ctx, org, opt)
	return args.Get(0).([]*github.Repository), args.Get(1).(*github.Response), args.Error(2)
}

func (m *MockGitHubClient) ListAllTopics(ctx context.Context, owner string, repo string) ([]string, *github.Response, error) {
	args := m.Called(ctx, owner, repo)
	return args.Get(0).([]string), args.Get(1).(*github.Response), args.Error(2)
}

func (m *MockGitHubClient) ReplaceAllTopics(ctx context.Context, owner string, repo string, topics []string) ([]string, *github.Response, error) {
	args := m.Called(ctx, owner, repo, topics)
	return args.Get(0).([]string), args.Get(1).(*github.Response), args.Error(2)
}
