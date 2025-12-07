// Data Types in Go
package main

import "fmt"

func main() {
	// String data type
	var str string = "Hello, Go!"
	fmt.Println("String:", str)
	// Integer data type
	var num int = 42
	fmt.Println("Integer:", num)
	// Float data type
	var floatNum float64 = 3.14
	fmt.Println("Float:", floatNum)
	// Boolean data type
	var isGoFun bool = true
	fmt.Println("Boolean:", isGoFun)
	// Complex data type
	var complexNum complex128 = complex(2, 3)
	fmt.Println("Complex:", complexNum)
	// Rune data type (represents a Unicode code point)
	var char rune = 'G'
	fmt.Println("Rune:", char)
	// Byte data type (alias for uint8)
	var byteVal byte = 255
	fmt.Println("Byte:", byteVal)
}
