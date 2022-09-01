package main

import "fmt"

//for is Go’s only looping construct
func main() {
	//The most basic type, with a single condition.
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	//A classic initial/condition/after for loop.
	for j := 7; j <= 9; j ++ {
		fmt.Println(j)
	}

	//for without a condition will loop repeatedly until you break out of the loop or return from the enclosing function.
	for {
		fmt.Println("loop")
		break
	}

	//You can also continue to the next iteration of the loop.
	for n := 0; n <= 5; n ++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

	//"after" is optional
	for i := 1; i < 5; {
		fmt.Println(i)
		i++
	}

	//"initial" is optional
	c := 0
	for ;c < 5; c ++ {
		fmt.Println(c)
	}

	//condition is optional
	for i := 0; ; i++ {
		fmt.Println(i)
		if (i == 5) {
			break
		}
	}
}
