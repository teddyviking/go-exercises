package accumulate

const testVersion = 1

func Accumulate(elements []string, operation func(string) string) []string {
	result := make([]string, len(elements))
	for i, element := range elements {
		result[i] = operation(element)
	}
	return result
}
