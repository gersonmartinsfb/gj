package gemini

import (
	"context"
	"log"

	"github.com/gersonmartinsfb/gj/config"

	"google.golang.org/genai"
)

type AdapterGemini struct {
	key   string
	model string
}

func NewAdapterGemini() *AdapterGemini {
	return &AdapterGemini{
		key:   config.Env.GeminiKey,
		model: "gemini-1.5-flash",
	}
}

func (a *AdapterGemini) GetContentFromPrompt(prompt string) (string, error) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  a.key,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}

	result, err := client.Models.GenerateContent(
		ctx,
		a.model,
		genai.Text(prompt),
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	if len(result.Candidates) == 0 {
		return "", nil
	}

	return result.Candidates[0].Content.Parts[0].Text, nil

}
