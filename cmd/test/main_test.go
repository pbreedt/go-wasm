package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestProcess(t *testing.T) {
	prompt := "prompt: "

	tests := []struct {
		fullText string
		command  string
		newInput string
	}{
		{"prompt: command", "command", ""},
		{"prompt: command with input", "command", "with input"},
		{"prompt: ", "", ""},
	}

	for _, tc := range tests {
		a := []string{prompt, tc.fullText}
		p, _, c, n := processTest(a)

		if p != prompt {
			t.Fatalf("prompt failed: expected: %s, got: %s", prompt, p)
		}

		if c != tc.command {
			t.Fatalf("command failes: expected: %s, got: %s", tc.command, c)
		}

		if n != tc.newInput {
			t.Fatalf("new input failes: expected: %s, got: %s", tc.newInput, n)
		}
	}
}

func processTest(args []string) (string, string, string, string) {
	prompt := args[0]
	text := args[1]

	lines := strings.Split(text, "\n")
	lastLine := lines[len(lines)-1]
	lastLineNoPrompt := lastLine[len(prompt):]
	command := lastLineNoPrompt
	newInput := ""
	if strings.Contains(lastLineNoPrompt, " ") {
		command = lastLineNoPrompt[:strings.Index(lastLineNoPrompt, " ")]
		newInput = lastLineNoPrompt[strings.Index(lastLineNoPrompt, " ")+1:]
	}

	fmt.Printf("wasm prompt:'%s' | last-line-no-prompt:'%s' | command:'%s' | input:'%s'\n", prompt, lastLineNoPrompt, command, newInput)

	return prompt, lastLineNoPrompt, command, newInput
}
