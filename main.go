package main

// Import packages
import (
	"fmt"
	"os"
	"io/ioutil"
	"strings"
)

// Main function
func main() {
	key := os.Args[1]
	newValue := os.Args[2]

	file, err := ioutil.ReadFile(".env")
	checkError(err)

	content := string(file)
	i := strings.Index(content, key+"=")
	fullPhrase := discoverFullPhrase(content, i)
	replacement := key + "=" + newValue
	newContent := strings.Replace(content, fullPhrase, replacement, -1)
	writeFile(newContent)

	fmt.Println("Done! :D")
}

// Check and show error message if .env file doesn't exist.
func checkError(err error) {
	if err != nil {
		panic("Sorry! .env file doesn't exist")
	}
}

// Write key and value in .env file.
func writeFile(content string) {
	dumped := []byte(content)
	err := ioutil.WriteFile(".env", dumped, 0644)
	checkError(err)
}

// Search for key and value string in file.
func discoverFullPhrase(content string, start int) string {
	var result string

	for i := start; i < len(content); i++ {
		var string = string(content[i])
		if string == "\n" {
			break
		}
		result += string
	}
	return result
}
