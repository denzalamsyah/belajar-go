package Looping

import "strings"

func Reversed(str string) string{

	var reverse string
	data := strings.Split(str, " ")

	for i := len(data)-1; i >= 0; i--{

		for j := len(data[i])-1; j >= 1; j--{
			reverse += string(data[i][j]) + "_"
		}
		reverse += string(data[i][0])
		if i != 0{
			reverse += " "
		}
	}
	return reverse
}