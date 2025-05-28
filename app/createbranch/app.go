package createbranch

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/gersonmartinsfb/gj/adapters/jira"
	"github.com/gersonmartinsfb/gj/config"
)

type CreateBranch struct {
	adapter      *jira.Request
	createBranch bool
	maxLength    int
}

func NewCreateBranch() *CreateBranch {
	config := config.Env
	return &CreateBranch{
		adapter:      jira.NewRequest(config.JiraDomain, config.JiraUser, config.JiraToken, config.JiraIssuePrefix),
		createBranch: config.CreateBranch,
		maxLength:    config.MaxLength,
	}
}

func (cb *CreateBranch) CreateBranch(issueType string, issueID string) error {
	branchName, err := cb.getBranchName(issueType, issueID)
	if err != nil {
		return fmt.Errorf("error getting branch name: %w", err)
	}

	if cb.createBranch {
		cmd := exec.Command("git", "checkout", "-b", branchName)
		if err := cmd.Run(); err != nil {
			return err
		}
	} else {
		fmt.Println("git checkout -b", branchName)
	}

	return nil
}

func (cb *CreateBranch) getBranchName(issueType string, issueID string) (string, error) {
	description, err := cb.getIssueDescription(issueID)
	if err != nil {
		return "", err
	}

	branchName := cb.removeCharacters(description)

	return fmt.Sprintf("%s/%s-%s/%s", issueType, config.Env.JiraIssuePrefix, issueID, branchName), nil
}

func (cb *CreateBranch) getIssueDescription(issueID string) (string, error) {
	issueResponse, err := cb.adapter.GetIssueDetails(issueID)
	if err != nil {
		return "", err
	}
	return issueResponse.GetSummary(), nil
}

func (cb *CreateBranch) removeCharacters(s string) string {
	// Convert to lowercase
	s = strings.ToLower(s)

	// Replace spaces with dashes
	s = strings.ReplaceAll(s, " ", "-")

	// Remove all characters except a-z, 0-9, and dashes
	result := ""
	for _, char := range s {
		if (char >= 'a' && char <= 'z') || (char >= '0' && char <= '9') || char == '-' {
			result += string(char)
		}
	}

	// Truncate to maxLength
	if len(result) > cb.maxLength {
		result = result[:cb.maxLength]
	}

	return result
}
