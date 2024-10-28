package main

// fmt stands for the Format package. This package is all about formatting input and output.
// bufio = to get user input.
// strconv = to convert string to another type.
// strings = to manipulating strings.
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// every go file must have a main funcation same as c programming.
func main() {
	// helloWorld()
	// typeDeclaration()
	// variable()
	// userInput()
	// addToNumber()
	pointers()
}

func helloWorld() {
	fmt.Println("Hello, world!")
}

func typeDeclaration() {
	// define a alias for a new type
	type name string
	type age int

	// declare variables using that alias
	var name1 name = "surya"
	var age1 age = 24
	fmt.Println(name1, age1)

	var name2 name = "karmakar"
	var age2 age = 28
	fmt.Println(name2, age2)
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

func userInput() {
	// Stdin = standard input
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")

	// string data, error object. if you want to ignore error object variable then use underscore.
	input, _ := reader.ReadString('\n')
	fmt.Println("You entered:", input)

	// convert string to float number
	fmt.Print("Enter a number: ")
	floatInput, _ := reader.ReadString('\n')
	// strings.TrimSpace = if user enter space end of the number
	aFloat, err := strconv.ParseFloat(strings.TrimSpace(floatInput), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Value of number: ", aFloat)
	}

	// conver string to int
	fmt.Print("Enter a int: ")
	intInput, _ := reader.ReadString('\n')
	aNumber, err := strconv.ParseInt(strings.TrimSpace(intInput), 0, 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Value of number: ", aNumber)
	}
}

func addToNumber() {
	var i int = 5
	var f float64 = 10.5

	// here if i try to add int + float then its throw an missmatched type error.
	sum := float64(i) + f
	fmt.Print("Sum: ", sum)
}

func pointers() {

}
