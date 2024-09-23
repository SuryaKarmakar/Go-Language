package main

// fmt stands for the Format package. This package is all about formatting input and output.
import (
	"fmt"
)

// every go file must have a main funcation same as c programming.
func main() {
	// helloWorld()
	// variable()

}

func helloWorld() {
	fmt.Println("Hello world !")
}

func variable() {
	// # explicit typing
	// declare and initialize variable
	var str string = "this is a string"
	fmt.Println(str)

	var i int = 10
	fmt.Println(i)

	var defaultInt int
	fmt.Println(defaultInt) // if you declare a int but not assign any value then its store 0 by default.

	// # implicit typing
	// in implicit typing the variable declear the initilize types.
	var autoString = "compiler automatic apply a data type"
	fmt.Println(autoString)

	var autoInt = 20
	fmt.Println(autoInt)

	// :=
	// you can use := operator to declare and initialize variable without using var keyword
	// := operator only works inside the funcation. if you declare variable outside of the funcation then you have to use var keyword
	myStr := "this is my string using := operator"
	fmt.Println(myStr)

	// to declare constant
	const pi float32 = 3.14
	fmt.Println(pi)
}
