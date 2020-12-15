// Package rlang provides functions for github repository language scanning.
package rlang

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/google/go-github/github"
	"github.com/siuyin/dflt"
	"golang.org/x/oauth2"
)

// NewClient returns an authenticated github client.
// Requires TOKEN environment variable to be correctly set.
func NewClient() *github.Client {
	tok := dflt.EnvString("TOKEN", "1234....")
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: tok},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)
	return client
}

// Repos returns a list of repository objects via the github API.
func Repos(client *github.Client) []*github.Repository {
	maxNum, err := dflt.EnvInt("MAXNUM", 1000)
	if err != nil {
		log.Fatal(err)
	}
	opt := &github.RepositoryListByOrgOptions{
		ListOptions: github.ListOptions{PerPage: maxNum},
	}
	org := dflt.EnvString("ORG", "someOrg")
	repos, _, err := client.Repositories.ListByOrg(context.Background(), org, opt)
	if err != nil {
		log.Fatalf("could not get repos: %v", err)
	}
	return repos
}

// CSVReport produces a report from repository list.
func CSVReport(repos []*github.Repository) {
	cw := csv.NewWriter(os.Stdout)
	defer cw.Flush()

	writeCSVHeader(cw)
	for _, v := range repos {
		lang := ""
		if v.Language != nil {
			lang = *v.Language
		}
		cw.Write([]string{*v.Name, fmt.Sprintf("%v", *v.ID), lang})
	}

}
func writeCSVHeader(w *csv.Writer) {
	w.Write([]string{"Name", "ID", "Language"})
}
