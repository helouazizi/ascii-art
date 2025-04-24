package functions

import (
	"fmt"
	"os"
	"strings"
)

// this func just to read the banner file
func ReadFile(filename string) []string {
	data, err := os.ReadFile(filename)
	// handle err
	if err != nil {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --output=<fileName.txt> something standard")
		os.Exit(0)
	}
	// handling if the banner file was writenf by windows
	stringdata := string(data)
	stringdata = strings.ReplaceAll(stringdata, "\r\n", "\n")

	result := strings.Split(stringdata, "\n")
	return result
}

// this is the the traitment functions
func TraitmentData(text []byte, arg, resultFile string) {
	// cheeck if the char is in range or not if
	for _, char := range arg {
		if char < 32 || char > 126 {
			fmt.Println("Ereur : one of this carcters is not in range")
			return
		}
	}
	result := ""

	arrData := strings.Split(string(text), "\n")

	words := strings.Split(arg, `\n`)

	/////////////////////////////////
	// this part just for newlines
	count := 0
	for _, test := range words {
		if test == "" {
			count++
		}
	}
	// in case the data is all new line
	if count == (len(arg)/2)+1 {
		for i := 0; i < (len(arg) / 2); i++ {
			result += "\n"
		}
	} else {
		result = Final_result(arrData, words)
	}
	if resultFile != "" {
		os.WriteFile(resultFile, []byte(result), 0o777)

	} else {
		fmt.Print(result)
	}

}

// traitment the data if it have charachters
func Final_result(arrData, words []string) string {
	result := ""
	for k := 0; k < len(words); k++ {
		if words[k] == "" {
			result += "\n"
			continue
		}
		for i := 0; i < 8; i++ {
			for j := 0; j < len(words[k]); j++ {
				Ascii := (int(words[k][j] - 32))

				start := Ascii*8 + Ascii + 1 + i
				if words[k][j] < 32 || words[k][j] > 126 {
					fmt.Println(" error : one of this charachter not in range ")
					os.Exit(0)
				} else {
					result += arrData[start]
				}

			}
			result += "\n"
		}
	}

	return result
}

// SafeFile checks whether a file path is considered safe for writing
func SafeFile(path string) bool {

	// Define restricted directories
	restrictedDirs := []string{
		"./banners/",
		"./functions/",
		"./parser/",
	}

	// Check if the absolute path is within any restricted directories
	for _, dir := range restrictedDirs {
		if strings.HasPrefix(path, dir) {
			return false // Path is inside a restricted directory
		}
	}
	if path == "main.go" || path == "go.mod" {
		return false
	}
	return true
}
