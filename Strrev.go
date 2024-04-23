package piscine

import "fmt"



func ReverseString(s string)string{
	runes:=[]rune(s)

	for i:= len(runes)-1;i>=0;i--{
		fmt.Println(string(runes[i]))
	}
return string(runes)

}