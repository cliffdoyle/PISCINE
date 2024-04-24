package piscine

func IterativeFactorial(num int) int {
	if num < 0 {
		return 0
		}
	if num > 63 {
		return 0
		}
	result := 1

	for i := 1; i < num+1; i++ {
		result = result * i
	}
	return result
}
