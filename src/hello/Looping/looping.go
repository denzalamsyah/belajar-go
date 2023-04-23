package Looping

import "fmt"

func Loopings(n int) int {
	for i := 1; i <= n; i++ {
		for j := 2; j <= 3; j++ {
			fmt.Printf("Hello Go! i: %d, j: %d\n", i, j)
		}
	}
	return n

	// i := 1
	// for i < 5 {
	// 	fmt.Println("ini angka ke - ", i)
	// 	i++
	// }

	//Empity for statement

	// i := 1
	// for {
	// 	if i >= n {
	// 		break
	// 	}

	// 	fmt.Println("ini angka ke - ", i)
	// 	i++

	// }

}