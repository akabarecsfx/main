package main

import (
	"fmt"
)

func isVokal(character rune) {
	switch character {
	case 'a', 'e', 'i', 'o', 'u':
		fmt.Printf(" %c adalah vokal\n", character)
	default:
		fmt.Printf(" %c adalah konsonan\n", character)
	}
}
func main() {
	isVowel('e')
	isVowel('g')
}
