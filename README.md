# Github Repository Languages
github-repo-langs generates a report of languages used within
an organisation.

## Configuration and run
1. cp sample.env yourSecret.env

1. edit yourSecret.env accordingly

1. . yourSecret.env <-- it reads "dot yourSecret.env"

go run cmd/repolangs/main.go > yourReport.csv
