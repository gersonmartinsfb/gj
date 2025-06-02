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
		fmt.Fprintf(os.Stderr, "  type: The type of issue (bug, feat, hotfix, enhance). Could be the initial letter (b, f, h, e)\n")
		fmt.Fprintf(os.Stderr, "  issue-number: The issue number to create a branch for. Just the number. The suffix %s- will be added automatically.\n", config.Env.JiraIssuePrefix)

		fmt.Fprintf(os.Stderr, "\nExtra:\n")
		fmt.Fprintf(os.Stderr, "    feat -- Para feature\n")
		fmt.Fprintf(os.Stderr, "    hotfix -- Problema que precisa entrar como prioridade em prod ASAP\n")
		fmt.Fprintf(os.Stderr, "    bug -- Resolução de bug\n")
		fmt.Fprintf(os.Stderr, "    enhance -- Refactor / melhoria / otimização / lynt / resolução de ofensas\n")

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

	argType = returnFullName(argType)

	branch := createbranch.NewCreateBranch()
	err := branch.CreateBranch(argType, issueNumber)
	if err != nil {
		fmt.Printf("Error creating branch for issue %s: %v\n", issueNumber, err)
		os.Exit(1)
	}
}

func returnFullName(argType string) string {
	switch argType {
	case "b", "bug":
		return "bug"
	case "f", "feat":
		return "feat"
	case "h", "hotfix":
		return "hotfix"
	case "e", "enhance":
		return "enhance"
	default:
		fmt.Printf("Unknown type: %s\n", argType)
		fmt.Println("Valid types are: b/bug, f/feat, h/hotfix, e/enhance")
		os.Exit(1)
	}

	return ""
}
