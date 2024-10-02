# GitHub Repository Tagging Script
This script is designed to list and process repositories within a specified GitHub organization. It reads a JSON file containing keyword topics and uses the GitHub API to retrieve repository information.

## Prerequisites
- Go 1.16 or later
- A GitHub personal access token with the necessary permissions to list repositories in the organization

## Configuration
1. Create a JSON file containing keyword topics. The structure of the JSON file should match the expected format in the script.

2. Install dependencies
    ```
    go mod tidy
    ```

2. Set up environment variables for your GitHub personal access token and the organization name:
    ```
    export GITHUB_TOKEN=<personal-access-token>
    export GITHUB_ORG=<organization-name>
    ```

## Usage
Run the script using the following command:
```
go run main_tagging.go
```

## Code Explanation
The script performs the following steps:

1. Reads and unmarshals the JSON file containing keyword topics.

2. Sets up options for listing repositories in the specified GitHub organization.

3. Retrieves all repositories from the organization, handling pagination.

4. (Commented out) Processes each repository. This section is commented out to avoid updating all repositories during testing.

