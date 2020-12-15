package main

import (
	"github.com/siuyin/github-repo-langs/rlang"
)

func main() {
	client := rlang.NewClient()
	repos := rlang.Repos(client)
	rlang.CSVReport(repos)
}
