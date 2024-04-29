package piscine


func Fibonacci(num int)int{
	if num < 0{
		return -1
	}

	return Fibonacci(num-1)+Fibonacci(num-2)
}