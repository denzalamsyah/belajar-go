package Condition

import "fmt"

func Conditions() {
	// var score, minimal = 80, 61

	// var result = score > 70 && minimal > 60
	// fmt.Println(result)

	// If else

	// i := 5

	// if i > 1{
	// 	fmt.Println("Hello Word") //yang pertama yang akan dieksekusi
	// }else if i > 2{
	// 	fmt.Println("Hello word word")
	// }else{
	// 	fmt.Println("hello")
	// }

	//Swit case

	day := 1
	switch {
	case day < 2:
		fmt.Println("hello word")
		fallthrough
	case day < 3:
		fmt.Println("hello word word")
		fallthrough
	case day < 4:
		fmt.Println("Hello")
		fallthrough
	default:
		fmt.Println("hai")
	}
}