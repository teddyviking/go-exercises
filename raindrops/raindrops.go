package raindrops

import (
	"strconv"
)
const testVersion = 2

type Rule struct {
	Factor int
	Output string
}

var rules = [3]Rule{
	{3, "Pling"},
	{5, "Plang"},
	{7, "Plong"},
}

func Convert(number int) string {
	converted := ""
	for _, rule := range rules {
		if divisible(number, rule.Factor) {
			converted += rule.Output
		}
	}
	if converted == "" {
		converted = strconv.Itoa(number)
	}
	return converted
}

func divisible(number, target int) bool {
	return number % target == 0
}
