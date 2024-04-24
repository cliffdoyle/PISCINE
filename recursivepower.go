package piscine

func RecursivePower(num int, power int) int {
	if power < 0 {
		return 0
	}
	return num * RecursivePower(num, power-1)
}
