package main

import (
	_ "Hello/Condition"
	"Hello/Looping"
	_ "Hello/Looping"
	"fmt"
)




func main() {
	// Looping.Loopings(2)
	// Condition.Conditions()
	fmt.Println(Looping.Reversed("Hello Word"))
}