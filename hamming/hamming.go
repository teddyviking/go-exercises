package hamming

import (
	"errors"
	"strings"
	"fmt"
)

const testVersion = 5

func Distance(a, b string) (int, error) {
	result := 0
	lettersA := strings.Split(a, "")
	lettersB := strings.Split(b, "")

	if len(lettersA) != len(lettersB) {
		return result, errors.New("Invalid length of strings")
	}

	for index, letter := range lettersA  {
		if letter != lettersB[index] {
			result += 1
		}
	}
	fmt.Printf("Distance: %d\n", result)
	return result, nil
}
