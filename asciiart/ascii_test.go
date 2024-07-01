package ascii

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestAsciiArt(t *testing.T) {
	type parameters struct {
		words    []string
		contents []string
	}
	type testCases struct {
		name     string
		args     parameters
		expected string
	}
	// We read the bannerfile directly to save on space/make the code readable
	// otherwise we would have provided the contents in banner file as a []string.
	contents, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Print("Error reading from file", err)
	}
	expArt, err := os.ReadFile("expected.txt")
	if err != nil {
		fmt.Print("Error reading from file", err)
	}
	textArt := strings.Split(string(contents), "\n")
	tests := []testCases{
		{name: "test empty", args: parameters{words: []string{}, contents: textArt}, expected: ""},
		{name: "test newline(\n)", args: parameters{words: []string{""}, contents: textArt}, expected: ""},
		{name: "test non Ascii", args: parameters{words: []string{"你好"}, contents: textArt}, expected: "Error: Non printable ASCII character\n"},
		{name: "test Hello", args: parameters{words: []string{"Hello"}, contents: textArt}, expected: string(expArt)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := AsciiArt(tt.args.words, tt.args.contents)
			if actual != tt.expected {
				t.Errorf("Output: %s Expected: %s\n", actual, tt.expected)
			}
		})
	}
}
