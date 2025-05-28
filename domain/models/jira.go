package models

type Fields struct {
	Summary string `json:"summary"`
}

type JiraIssueResponse struct {
	Summary string `json:"summary"`
	Fields  Fields `json:"fields"`
}

func (j *JiraIssueResponse) GetSummary() string {
	if j.Fields.Summary != "" {
		return j.Fields.Summary
	}
	return j.Summary
}
