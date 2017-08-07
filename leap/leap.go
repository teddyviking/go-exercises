package leap

const testVersion = 3

func IsLeapYear(year int) bool {
	result := false

	commonRest := year % 4
	hundredRest := year % 100
	fourHundredRest := year % 400

	if fourHundredRest == 0 || (commonRest == 0 && hundredRest != 0) {
		result = true
	}

	return result
}
