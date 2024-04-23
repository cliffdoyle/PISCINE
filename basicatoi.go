package piscine

func BasicAtoi(s string) int {
	result := 0

	for _, char := range s {
		digits := int(char - '0')
		result = result*10 + digits
	}

	return result
}
