package lang

import (
	"golang_ot/lang/english"
	"golang_ot/lang/hindi"
	"strings"
)

var Messages = map[string]struct {
	General map[string]string
	Words   map[string]string
}{
	"English": {
		General: english.General,
		Words:   english.Words,
	},
	"Hindi": {
		General: hindi.General,
		Words:   hindi.Words,
	},
}

// GetMessage retrieves a translated message from the general category
func GetMessage(lang, key string) string {
	if data, exists := Messages[lang]; exists {
		if val, ok := data.General[key]; ok {
			return val
		}
	}
	return key // Return key if not found
}

// GetWord retrieves a translated word from the words category
func GetWord(lang, key string) string {
	if data, exists := Messages[lang]; exists {
		if val, ok := data.Words[key]; ok {
			return val
		}
	}
	return key // Return key if not found
}

// ReplacePlaceholders replaces placeholders (##) in a message with words
func ReplacePlaceholders(langCode string, message string, replacements ...string) string {
	for _, replacement := range replacements {
		replacementWord := GetWord(langCode, replacement)
		message = strings.Replace(message, "##", replacementWord, len(replacements))
	}
	return message
}
