package main

import "fmt"

//Go requires explicit returns, i.e. it won’t automatically return the value of the last expression.
func plus(a int, b int) int {
	return a + b
}

//When you have multiple consecutive parameters of the same type,
//you may omit the type name for the like-typed parameters up to the final parameter that declares the type.
func plusPlus(a, b, c int) int {
	return a + b + c
}

//The (int, int) in this function signature shows that the function returns 2 ints.
func vals() (int, int) {
	return 3, 7
}

//Here’s a function that will take an arbitrary number of ints as arguments.
//aka Variadic function
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0

	//Within the function, the type of nums is equivalent to []int.
	//We can call len(nums), iterate over it with range, etc.
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

//Call a function just as you’d expect, with name(args).
func main() {

	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)

	//Here we use the 2 different return values from the call with multiple assignment.
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	//If you only want a subset of the returned values, use the blank identifier _.
	_, c := vals()
	fmt.Println(c)

	//Variadic functions can be called in the usual way with individual arguments.
	sum(1, 2)
	sum(1, 2, 3)

	//If you already have multiple args in a slice, apply them to a variadic function using func(slice...) like this.
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
