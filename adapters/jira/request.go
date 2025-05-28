package jira

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gersonmartinsfb/gj/domain/models"
)

type Request struct {
	domain string
	user   string
	token  string
	prefix string
}

func NewRequest(domain, user, token, prefix string) *Request {
	return &Request{
		domain: domain,
		user:   user,
		token:  token,
		prefix: prefix,
	}
}

func (r *Request) GetIssueDetails(issueID string) (models.JiraIssueResponse, error) {
	url := "https://" + r.domain + "/rest/api/3/issue/" + r.prefix + "-" + issueID
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return models.JiraIssueResponse{}, err
	}
	req.SetBasicAuth(r.user, r.token)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return models.JiraIssueResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return models.JiraIssueResponse{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.JiraIssueResponse{}, err
	}

	var issueResponse models.JiraIssueResponse
	if err := json.Unmarshal(body, &issueResponse); err != nil {
		return models.JiraIssueResponse{}, err
	}

	return issueResponse, nil
}
