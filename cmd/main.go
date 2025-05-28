package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gersonmartinsfb/gj/app/createbranch"
	"github.com/gersonmartinsfb/gj/config"
)

func main() {
	config.Load()
	disableBranch := flag.Bool("disable-creation", false, "Disable branch creation and just print the command that would be run")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: gj [--disable-creation] <type> <issue-number>\n")
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "  type: The type of issue (bug, feat, hotfix, enhance)\n")
		fmt.Fprintf(os.Stderr, "  issue-number: The issue number to create a branch for. Just the number. The suffix %s- will be added automatically.\n", config.Env.JiraIssuePrefix)
	}

	flag.Parse()

	args := flag.Args()
	if len(args) < 2 {
		fmt.Println("Usage: gj [--disable-creation] <type> <issue-number>")
		os.Exit(1)
	}

	argType := args[0]
	issueNumber := args[1]
	config.Env.CreateBranch = !(*disableBranch)

	branch := createbranch.NewCreateBranch()
	err := branch.CreateBranch(argType, issueNumber)
	if err != nil {
		fmt.Printf("Error creating branch for issue %s: %v\n", issueNumber, err)
		os.Exit(1)
	}
}
