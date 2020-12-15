package main

import (
	"context"
	"fmt"
	"log"

	"github.com/siuyin/dflt"
	"golang.org/x/oauth2"
)

func main() {
	fmt.Println("Scanning Github repository languages")

	tok := dflt.EnvString("TOKEN", "1234....")
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: tok},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	// list all repositories for the authenticated user
	repos, _, err := client.Repositories.List(ctx, "", nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(repos)
}
