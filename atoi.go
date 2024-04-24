package piscine

func RealAtoi(s string) int {
	result := 0
	sign := 1
	signDetected := false

	for _, char := range s {
		if char == '-' || char == '+' {
			if signDetected {
				return 0
			}
			if char == '-' {
				sign = -1
			}
			signDetected = true
			continue
		}
		if char < '0' || char > '9' {
			return 0
		} else {
			digits := int(char - '0')
			result = result*10 + digits
		}

	}
	return result * sign
}
