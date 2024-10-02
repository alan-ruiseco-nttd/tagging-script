package main

import (
	"context"
	"testing"

	"github.com/google/go-github/v65/github"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAllRepositories(t *testing.T) {
	mockClient := new(MockGitHubClient)
	ctx := context.Background()
	org := "test-org"

	mockRepos := []*github.Repository{
		{Name: github.String("repo1")},
		{Name: github.String("repo2")},
	}

	mockClient.On("ListByOrg", ctx, org, mock.Anything).Return(mockRepos, &github.Response{}, nil)

	repos := getAllRepositories(ctx, mockClient, org)

	mockClient.AssertExpectations(t)
	assert.Equal(t, 2, len(repos))
	assert.Equal(t, "repo1", *repos[0].Name)
	assert.Equal(t, "repo2", *repos[1].Name)
}

func TestGetCurrentTopics(t *testing.T) {
	mockClient := new(MockGitHubClient)
	ctx := context.Background()
	org := "test-org"
	repo := "test-repo"

	mockTopics := []string{"azure", "aws"}

	mockClient.On("ListAllTopics", ctx, org, repo).Return(mockTopics, &github.Response{}, nil)

	topics := getCurrentTopics(ctx, mockClient, org, repo)

	mockClient.AssertExpectations(t)
	assert.Equal(t, 2, len(topics))
	assert.Equal(t, "azure", topics[0])
	assert.Equal(t, "aws", topics[1])
}

func TestUpdateCurrentTopics(t *testing.T) {
	keywordTopics := map[string]string{"azurerm-": "azure", "aws-": "aws"}
	currentTopics := []string{"azure", "aws"}
	repoName := "test-repo"

	topics, updated := updateCurrentTopics(currentTopics, keywordTopics, repoName)

	if updated {
		assert.True(t, updated)
	}
	assert.Equal(t, 2, len(topics))
	assert.Equal(t, "azure", topics[0])
}

func TestApplyNewTopics(t *testing.T) {
	mockClient := new(MockGitHubClient)
	ctx := context.Background()
	org := "test-org"
	repo := "test-repo"
	topics := []string{"azure", "aws"}

	mockClient.On("ReplaceAllTopics", ctx, org, repo, topics).Return([]string{"azure", "aws"}, &github.Response{}, nil)

	applyNewTopics(ctx, true, mockClient, org, repo, topics)

	mockClient.AssertExpectations(t)
}
