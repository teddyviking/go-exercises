package acronym

import (
	"strings"
)

const testVersion = 2

func Abbreviate(target string) string {
	var result string
	words := strings.FieldsFunc(target, split)
	for _, word := range words {
		word = capitalize(word)
		result += string(word[0])
		result += getUpperLetters(word)
	}

	return result
}

func capitalize(word string) string {
	return strings.Title(word)
}

func split(r rune) bool {
	return r == ' ' || r == '-'
}

func getUpperLetters(word string) string {
	var result string
	for i, letter := range word {
		if i == 0 {
			continue
		}
		stringed := string(letter)
		previous := string(word[i-1])

		if strings.ToUpper(stringed) == stringed && strings.ToUpper(previous) != previous && !strings.ContainsAny(stringed, ","){
			result = stringed
			break
		}
	}
	return result
}
