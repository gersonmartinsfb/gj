package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvT struct {
	JiraToken       string
	JiraUser        string
	JiraDomain      string
	JiraIssuePrefix string
	CreateBranch    bool
	MaxLength       int
}

var Env EnvT

func Load() {
	_ = godotenv.Load(".env")

	prefix := os.Getenv("GJ_ISSUE_PREFIX")
	if prefix == "" {
		log.Println("GJ_ISSUE_PREFIX not set, defaulting to 'QUANT'")
		prefix = "QUANT"
	}

	Env = EnvT{
		JiraToken:       mustGetEnv("GJ_TOKEN"),
		JiraUser:        mustGetEnv("GJ_USER"),
		JiraDomain:      mustGetEnv("GJ_DOMAIN"),
		JiraIssuePrefix: prefix,
		CreateBranch:    true,
		MaxLength:       50, // Default max length for branch names
	}
}

func mustGetEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Fatalf("missing required env var: %s", key)
	}
	return val
}
