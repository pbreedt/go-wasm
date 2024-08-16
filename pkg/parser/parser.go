package parser

import (
	"strings"
)

func ParseInput(prompt, text string) (string, string, string, string) {
	lines := strings.Split(text, "\n")
	lastLine := lines[len(lines)-1]
	lastLineNoPrompt := lastLine[len(prompt):]
	command := lastLineNoPrompt
	newInput := ""
	if strings.Contains(lastLineNoPrompt, " ") {
		command = lastLineNoPrompt[:strings.Index(lastLineNoPrompt, " ")]
		newInput = lastLineNoPrompt[strings.Index(lastLineNoPrompt, " ")+1:]
	}

	// fmt.Printf("wasm prompt:'%s' | last-line-no-prompt:'%s' | command:'%s' | input:'%s'\n", prompt, lastLineNoPrompt, command, newInput)

	return prompt, lastLineNoPrompt, command, newInput
}
