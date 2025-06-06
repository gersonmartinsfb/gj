package translate

import (
	"strings"

	"github.com/gersonmartinsfb/gj/adapters/gemini"
)

type Translate struct {
	adapter *gemini.AdapterGemini
}

func NewTranslate() *Translate {
	return &Translate{
		adapter: gemini.NewAdapterGemini(),
	}
}

func (t *Translate) TranslateText(text string) (string, error) {
	prompt := `Translate the following text to English. This text is a title for a Jira card.

Follow these rules:

If the text begins with a word enclosed in brackets (e.g., [TAG]), that word and the brackets must be ignored.
The final translated text cannot be more than 50 characters long.
If necessary, use common abbreviations to meet the character limit:
---
` + text
	translatedText, err := t.adapter.GetContentFromPrompt(prompt)
	if err != nil {
		return "", err
	}

	if translatedText == "" {
		return text, nil // If the translation is empty, return the original text
	}

	// Remove all tabs and new lines from translatedText
	translatedText = strings.ReplaceAll(translatedText, "\n", "")
	translatedText = strings.ReplaceAll(translatedText, "\r", "")
	translatedText = strings.ReplaceAll(translatedText, "\t", "")

	return translatedText, nil
}
