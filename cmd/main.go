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
	disableBranch := flag.Bool("disable-branch", false, "Disable branch creation")
	dFlag := flag.Bool("d", false, "Disable branch creation (shorthand)")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: gj [--disable-branch|-d] <type> <issue-number>\n")
		fmt.Fprintf(os.Stderr, "Options:\n")
		fmt.Fprintf(os.Stderr, "\ttype: The type of issue (bug, feat, hotfix, enhance)\n")
		fmt.Fprintf(os.Stderr, "\tissue-number: The issue number to create a branch for. Just the number. The suffix %s- will be added automatically.\n", config.Env.JiraIssuePrefix)
		flag.PrintDefaults()
	}

	flag.Parse()

	args := flag.Args()
	if len(args) < 2 {
		fmt.Println("Usage: gj [--disable-branch|-d] <type> <issue-number>")
		os.Exit(1)
	}

	argType := args[0]
	issueNumber := args[1]
	createBranch := !(*disableBranch || *dFlag)

	config.Env.CreateBranch = createBranch

	branch := createbranch.NewCreateBranch()
	err := branch.CreateBranch(argType, issueNumber)
	if err != nil {
		fmt.Printf("Error creating branch for issue %s: %v\n", issueNumber, err)
		os.Exit(1)
	}
}
