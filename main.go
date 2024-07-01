package main

import (
	"fmt"
	"os"
	"strings"

	ascii "ascii/asciiart"
)

func main() {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("Usage: go run . [STRING] [BANNER]\n\nEX.: go run . something standard")
		return
	}
	// Grab string to generate Ascii represatantion.
	inputText := os.Args[1]

	switch inputText {
	case "":
		return
	case "\\a", "\\0", "\\f", "\\v", "\\r":
		fmt.Println("Error: Non printable character", inputText)
		return
	}

	inputText = strings.ReplaceAll(inputText, "\\t", "    ")
	inputText = strings.ReplaceAll(inputText, "\\b", "\b")
	inputText = strings.ReplaceAll(inputText, "\\n", "\n")
	// Logic process for handlng the backspace.
	for i := 0; i < len(inputText); i++ {
		indexB := strings.Index(inputText, "\b")
		if indexB > 0 {
			inputText = inputText[:indexB-1] + inputText[indexB+1:]
		}
	}
	// Split our input text to a string slice and separate with a newline.
	words := strings.Split(inputText, "\n")

	// setting the bannerfile to be used according to user input.
	banner := "standard"

	if len(os.Args) == 3 {
		banner = strings.ToLower(os.Args[2])
	}
	// Check if the banner has an extension
	if strings.Contains(banner, ".") {
		// Check if the extension is not .txt
		if !strings.HasSuffix(banner, ".txt") {
			fmt.Println("Error: Required format: banner.txt")
			return
		}
	} else {
		// If no extension, add .txt
		banner = banner + ".txt"
	}

	// Convert to lowercase
	banner = strings.ToLower(banner)

	bannerFile := banner

	// Read the contents of banner file.
	bannerText, err := os.ReadFile(bannerFile)
	if err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}
	// Confirm file information.
	fileInfo, err := os.Stat(bannerFile)
	if err != nil {
		fmt.Println("Error reading file information", err)
		return
	}
	fileSize := fileInfo.Size()

	if fileSize == 6623 || fileSize == 4702 || fileSize == 7462 || fileSize == 4496 {
		// Split the content to a string slice and separate with newline.
		contents := strings.Split(string(bannerText), "\n")

		fmt.Print(ascii.AsciiArt(words, contents))

	} else {
		fmt.Println("Error with the file size", fileSize)
		return
	}
}
