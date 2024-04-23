package piscine

import "fmt"

func Printstr(s string) {
	for i := 0; i <= len(s); i++ {
		fmt.Println(s[i])
	}
}
