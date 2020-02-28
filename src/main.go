package main

import "fmt"

func main() {
	fmt.Println("Welcome to the beggining of my code.")

	a := 10

	fmt.Println("My number", a, "has been selected.")

	module_1()

	b := 20 - 10

	fmt.Println("My second selected number is", b)

	module_2()

	c := 1000

	fmt.Println("And finaly", c, "is my last number")

	module_3()
}

func module_1() {

	fmt.Println("This is the first module of my code.")
}
func module_2() {

	fmt.Println("This module will be shown in the middle of my code.")
}

func module_3() {

	fmt.Println("This is the end of my code.")
}