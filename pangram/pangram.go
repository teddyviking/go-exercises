package pangram

import (
	"strings"
)
const testVersion = 1
const alphabet = "abcdefghijklmnopqrstuvwxyz"

func IsPangram(sentence string) bool {
	lowered := strings.ToLower(sentence)
	for _, runeValue := range alphabet {
		if !strings.ContainsRune(lowered, runeValue) {
			return false
		}
	}
	return true
}
