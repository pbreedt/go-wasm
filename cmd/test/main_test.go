package main

import (
	"testing"

	"github.com/pbreedt/go-wasm/pkg/parser"
)

func TestProcess(t *testing.T) {
	prompt := "prompt: "

	tests := []struct {
		name     string
		fullText string
		command  string
		newInput string
	}{
		{"single command", "prompt: command", "command", ""},
		{"command with input", "prompt: command with input", "command", "with input"},
		{"no command", "prompt: ", "", ""},
	}

	for _, tc := range tests {
		t.Log("testing:", tc.name)
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

	return parser.ParseInput(prompt, text)
}
