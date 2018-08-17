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
	args := len(os.Args)

	switch args {
	case 2:
		getSomething(os.Args[1])

	case 3:
		setSomething(os.Args[1], os.Args[2])

	default:
		showHelp()
	}

	os.Exit(1)

}

// show help when unsupported number of arguments are given
func showHelp() {
	fmt.Println("App works with one or two number of arguments:")
	fmt.Println("- Use one argument to read something.")
	fmt.Println("- Use two arguments to set something.")
	fmt.Println("Pretty clear, hah?")
}

// get something from the .env file
func getSomething(key string) {
	content := readEnvFile()
	i := strings.Index(content, key+"=")

	if i < 0 {
		fmt.Println("Not Defined!")
		return
	}

	fullPhrase := discoverFullPhrase(content, i)
	parts := strings.Split(fullPhrase, "=")

	fmt.Println(parts[1])
}

// set something in the .env file
func setSomething(key string, newValue string) {
	key = strings.ToUpper(key)
	content := readEnvFile()

	i := strings.Index(content, key+"=")

	if i < 0 {
		setSomethingNew(content, key, newValue)
		return
	}

	setSomethingExisting(content, i, key, newValue)
}

// set some new key in the .env file
func setSomethingNew(content string, key string, newValue string) {
	newContent := content + "\n" + key + "=" + newValue
	writeFile(newContent)

	fmt.Println("A new", key, "has been inserted with value:", newValue)
}

// set some new key in the .env file
func setSomethingExisting(content string, index int, key string, newValue string) {
	fullPhrase := discoverFullPhrase(content, index)
	replacement := key + "=" + newValue
	newContent := strings.Replace(content, fullPhrase, replacement, -1)
	writeFile(newContent)

	fmt.Println("Value for" , key , "has been replaced with:" , newValue)

}

// try to open the file, returning a convenient message if failed.
func readEnvFile() string {
	file, err := ioutil.ReadFile(".env")

	if err != nil {
		fmt.Println("Error Occurred: Couldn't find a `.env` file in this route.")
		os.Exit(2)
	}

	return string(file)
}

// Write key and value in .env file.
func writeFile(content string) {
	dumped := []byte(content)
	err := ioutil.WriteFile(".env", dumped, 0644)

	if err != nil {
		fmt.Println("Error Occurred: Couldn't write the`.env` file in this route.")
		os.Exit(2)
	}
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
