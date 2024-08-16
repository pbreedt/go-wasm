package main

import (
	_ "embed"
	"fmt"
	"io"
	"net/http"
	"strings"
	"syscall/js"

	"github.com/pbreedt/go-wasm/pkg/parser"
)

//go:embed main.go
var source string

func main() {
	fmt.Println("Go Web Assembly")
	js.Global().Set("execGo", exec())
	js.Global().Set("execGoUbuntu", execUbuntu())
	<-make(chan struct{})
}

func exec() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 2 {
			return "Invalid no of arguments passed"
		}

		prompt := args[0].String()
		text := args[1].String()

		_, _, command, newInput := parser.ParseInput(prompt, text)

		return process(text, prompt, command, newInput)
	})
}

func process(originalText, prompt, command, newInput string) string {
	switch command {
	case "clear":
		return prompt
	case "profile":
		return addNewInputToEnd(originalText, profile(), prompt)
	case "source":
		return addNewInputToEnd(originalText, source, prompt)
	case "upper":
		return addNewInputToEnd(originalText, toUpper(newInput), prompt)
	default:
		return addNewInputToEnd(originalText, help(), prompt)
	}
}

func profile() string {
	return `Profile: Petrus Breedt
    - education : B.Sc Information Technology (University of Johannesburg, South Africa)
    - job title : Software Architect
    - skills    : Software development (Go, Java)
                  Requirements Analysis
                  Solution Design
                  Team lead & mentor`
}

func addNewInputToEnd(originalText, processedText, prompt string) string {
	return originalText + "\n" + processedText + "\n" + prompt
}

func toUpper(input string) string {
	return strings.ToUpper(input)
}

func help() string {
	return `
Usage: command [input...]

commands:
	help    : displays this message
	clear   : clears the screen
	upper   : returns uppercase of the input
	profile : displays user profile
	source  : displays site's Go source code`
}

func readFile(filename string) (string, error) {
	resp, err := http.Get("http://localhost:9090/" + filename)
	if err != nil {
		return "", fmt.Errorf("GET error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Status error: %v", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("Read body: %v", err)
	}

	return string(data), nil
}

// Ubuntu sets prompt in HTML
func execUbuntu() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 1 {
			return "Invalid no of arguments passed"
		}

		text := args[0].String()
		_, _, command, newInput := parser.ParseInput("", text)

		return process("", "", command, newInput)
	})
}
