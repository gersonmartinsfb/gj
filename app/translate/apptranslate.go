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

1. If the text begins with a word enclosed in brackets (e.g., [TAG]), that word and the brackets must be ignored.
2. The final translated text cannot be more than 100 characters long.
3. If necessary, use common abbreviations to meet the character limit
4. The translated text should be concise and clear, suitable for a Jira card title
5. Do not include any additional text or explanations, just the translated text
6. Do not include any special characters, emojis, or punctuation marks in the translated text 
7. Do not include any new lines, tabs, or extra spaces in the translated text
8. Try to translate all the text, but if it is not possible to translate all the text, just translate the most important part of the text
9. Do not replace 'and' word by '&' or 'and' by 'n', just translate the text as it is
10. If the text is already in English, return it as is
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
