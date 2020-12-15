package main

import (
	"context"
	"fmt"
	"log"

	"github.com/google/go-github/v33/github"
	"github.com/siuyin/dflt"
	"golang.org/x/oauth2"
)

func main() {
	fmt.Println("Scanning Github repository languages")

	client := getClient()

	repos := getRepos(client)

	genCSVReport(repos)
}

func getClient() *github.Client {
	tok := dflt.EnvString("TOKEN", "1234....")
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: tok},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)
	return client
}

func getRepos(client *github.Client) []*github.Repository {
	maxNum, err := dflt.EnvInt("MAXNUM", 1000)
	if err != nil {
		log.Fatal(err)
	}
	opt := &github.RepositoryListByOrgOptions{
		ListOptions: github.ListOptions{PerPage: maxNum},
	}
	repos, _, err := client.Repositories.ListByOrg(context.Background(), "xendit", opt)
	if err != nil {
		log.Fatalf("could not get repos: %v", err)
	}
	return repos
}

func genCSVReport(repos []*github.Repository) {
	for _, v := range repos {
		lang := ""
		if v.Language != nil {
			lang = *v.Language
		}
		fmt.Println(*v.Name, *v.ID, lang)
	}
}
