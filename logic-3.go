package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Masukkan string kurung: ")
	scanner.Scan()
	input := scanner.Text()
	fmt.Println(isValidString(input))
}

func isValidString(s string) bool {
	stack := []rune{}
	pairs := map[rune]rune{'<': '>', '{': '}', '[': ']'}

	for _, char := range s {
		switch char {
		case '<', '{', '[':
			stack = append(stack, char)
		case '>', '}', ']':
			if len(stack) == 0 {
				return false
			}
			last := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if pairs[last] != char {
				return false
			}
		}
	}
	return len(stack) == 0
}
