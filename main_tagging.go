package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"strings"

	"github.com/google/go-github/v65/github"
	"golang.org/x/oauth2"
)

func main() {
	ctx := context.Background()
	client := newGitHubClient(ctx)
	org := os.Getenv("GITHUB_ORG")

	allRepos := getAllRepositories(ctx, client, org)
	keywordTopics := getKeyWords()

	// Commented out to avoid updating all repos while testing
	for _, repo := range allRepos {
		repoName := *repo.Name

		// Test repos only
		// testRepos := []string{"tf-azurerm-module_reference-key_vault", "tf-aws-module_primitive-firehose_delivery_stream"}
		// for _, repoName := range testRepos {
		topics := getCurrentTopics(ctx, client, org, repoName)

		topics, updated := updateCurrentTopics(topics, keywordTopics, repoName)

		applyNewTopics(ctx, updated, client, org, repoName, topics)

	}
}

func newGitHubClient(ctx context.Context) *github.Client {
	token := os.Getenv("GITHUB_TOKEN")
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	return client
}

func getAllRepositories(ctx context.Context, client *github.Client, org string) []*github.Repository {
	opt := &github.RepositoryListByOrgOptions{
		ListOptions: github.ListOptions{PerPage: 100},
	}

	var allRepos []*github.Repository
	for {
		repos, resp, err := client.Repositories.ListByOrg(ctx, org, opt)
		if err != nil {
			log.Fatalf("Error listing repositories: %v", err)
		}
		allRepos = append(allRepos, repos...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}
	return allRepos
}

func getKeyWords() map[string]string {
	keywordsFile, err := os.ReadFile("topics.json")
	if err != nil {
		log.Fatalf("Error reading keywords JSON file: %v", err)
	}

	var keywordTopics map[string]string
	err = json.Unmarshal(keywordsFile, &keywordTopics)
	if err != nil {
		log.Fatalf("Error parsing keywords JSON file: %v", err)
	}
	return keywordTopics
}

func getCurrentTopics(ctx context.Context, client *github.Client, org string, repo string) []string {
	topics, _, err := client.Repositories.ListAllTopics(ctx, org, repo)
	if err != nil {
		log.Fatalf("Error listing topics for repo %s: %v", repo, err)
	}
	return topics
}

func updateCurrentTopics(topics []string, keywordTopics map[string]string, repoName string) ([]string, bool) {
	currentTopics := make(map[string]bool)
	for _, topic := range topics {
		currentTopics[topic] = true
	}

	updated := false
	for keyword, topic := range keywordTopics {
		if strings.Contains(repoName, keyword) && !currentTopics[topic] {
			topics = append(topics, topic)
			updated = true
		}
	}
	return topics, updated
}

func applyNewTopics(ctx context.Context, updated bool, client *github.Client, org string, repoName string, topics []string) {
	if updated {
		_, _, err := client.Repositories.ReplaceAllTopics(ctx, org, repoName, topics)
		if err != nil {
			log.Fatalf("Error updating topics for repo %s: %v", repoName, err)
		}
		log.Printf("Updated topics for repo %s", repoName)
	} else {
		log.Printf("No new topics to add for repo %s", repoName)
	}
}
