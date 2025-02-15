package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter length string array: ")
	scanner.Scan()
	n, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println(err)
		return
	}

	words := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Printf("Enter string n-%d: ", i+1)
		scanner.Scan()
		words[i] = scanner.Text()
	}

	result := findMatchStrings(words)
	fmt.Println(result)

}

func findMatchStrings(words []string) interface{} {
	seen := make(map[string][]int)
	firstDuplicate := ""

	for i, word := range words {
		lowerWord := strings.ToLower(word)
		seen[lowerWord] = append(seen[lowerWord], i+1)

		//check duplicate string
		if len(seen[lowerWord]) == 2 && firstDuplicate == "" {
			firstDuplicate = lowerWord
		}
	}

	if firstDuplicate != "" {
		return seen[firstDuplicate]
	}

	return false
}
