package main

import "fmt"

func main() {
	fmt.Println("Welcome to if else")

	count := 10
	var result string

	//part 1
	if count < 12 {
		result = "Less count"
	} else {
		result = "Max count"
	}
	fmt.Println(result)

	//part 2
	if num := 3; num < 3 {
		fmt.Println("Number is lessthan 3")
	} else {
		fmt.Println("Number is greate than 3")
	}

}
