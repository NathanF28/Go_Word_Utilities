package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func freq_count(sentence string) map[string]int {
	ans := make(map[string]int)
	temp := ""
	for _, v := range sentence {
		if v == ' ' || !unicode.IsLetter(v) {
			if len(temp) > 0 {
				ans[temp]++
			}
			temp = ""
		} else {
			temp += string(unicode.ToLower(v))
		}
	}
	if len(temp) > 0 {
		ans[temp]++
	}
	return ans
}

func is_palindrome(word string) string {
	front := ""
	for _, v := range word {
		if unicode.IsLetter(v) {
			front += string(unicode.ToLower(v))
		}
	}
	l := 0
	r := len(front) - 1
	for l < r {
		if front[l] != front[r] {
			return fmt.Sprintf("%v is not a palindrome", word)
		}
	}
	return fmt.Sprintf("%v is a palindrome", word)
}

func getStringInput(instruction string) string {
	fmt.Printf("\n\x1b[36m%v\x1b[0m", instruction)
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func getIntInput(instruction string) int {
	for {
		fmt.Printf("\n\x1b[36m%v\x1b[0m", instruction)
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		num, err := strconv.Atoi(text)
		if err == nil {
			return num
		}
		fmt.Println("\x1b[31mInvalid input, please enter a number.\x1b[0m")
	}
}

func prompt() {
	for {
		fmt.Println("\n\x1b[33m==============================\x1b[0m")
		fmt.Println("\x1b[33m1. Count frequency of words in a sentence\x1b[0m")
		fmt.Println("\x1b[33m2. Check if a word is a palindrome\x1b[0m")
		fmt.Println("\x1b[33m3. Exit program\x1b[0m")
		fmt.Println("\x1b[33m==============================\x1b[0m")

		num := getIntInput("Enter choice: ")

		switch num {
		case 1:
			sentence := getStringInput("Enter string: ")
			fmt.Println("\nResult:")
			fmt.Println("---------------------------")
			for word, count := range freq_count(sentence) {
				fmt.Printf("%-15s : %d\n", word, count)
			}
			fmt.Println("---------------------------")
		case 2:
			fmt.Print(is_palindrome(getStringInput("Enter string: ")))
		case 3:
			fmt.Println("\n\x1b[32mExiting program. Goodbye!\x1b[0m")
			return
		default:
			fmt.Println("\x1b[31mInvalid input\x1b[0m")
		}
	}

}
func main() {
	fmt.Println("\n\x1b[35m**************** Welcome to Utilities ****************\x1b[0m")
	prompt()
}
