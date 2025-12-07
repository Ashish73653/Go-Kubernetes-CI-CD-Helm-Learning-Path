package main

import "fmt"

func main() {
	var name string = "Ashish"
	fmt.Println(name)
	var val int = 7
	fmt.Println(val)
	var flag bool = true
	fmt.Println(flag)
	LastName := "Singh"
	fmt.Println(LastName)
	fmt.Println("Hi," + name + " " + LastName + " here." + " My Lucky number is: " + fmt.Sprint(val))
	// why sprint? because we cannot concatenate string with int directly

	// Multiple variable declaration with different types (string and int)
	// Go allows declaring multiple variables of different types in a single statement
	// var a, b = 6, "Hello"

	// Short declaration syntax with multiple variables of different types
	// The := operator infers types automatically from the assigned values
	// c, d := 7, "World!"

	// variables declared collectively
	var (
		c string = "Hello"
		d int    = 7
		e bool   = false
	)
	fmt.Println(c, d, e)

	// Constants in Go
	const Pi = 3.14
	const Greeting = "Hello, World!"
	const IsGoFun = true
	const MaxUsers = 1000
	const AppName = "GoApp"
	const Version = 1.0
	const DebugMode = false
	// types of constants
	const (
		ConstString  string  = "Constant String"
		ConstInt     int     = 42
		ConstFloat   float64 = 3.14159
		ConstBoolean bool    = true
	)
}
