package main

import "fmt"

//Variables are explicitly declared and used by the compiler to e.g. check type-correctness of function calls.
func main() {

	//var declares 1 or more variables.
	//Go will infer the type of initialized variables.
	var a = "initial"
	fmt.Println(a)

	//You can declare multiple variables at once.
	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	//Variables declared without a corresponding initialization are zero-valued. For example, the zero value for an int is 0.
	var e int
	fmt.Println(e)

	//The := syntax is shorthand for declaring and initializing a variable, e.g. for var f string = "apple" in this case.
	f := "apple"
	fmt.Println(f)

	//var g - error, expected type (can't infer type without initialization)
	//var x1 int, x2 int = 1, 2 - compile error. I.e. you declare type after variable names

	g := 1
	//g = "Test" - error, can't change variable type
	fmt.Println(g)
}