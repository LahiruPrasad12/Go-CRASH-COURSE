package main

import "fmt"

func main() {

	fmt.Println("welcome to go lang function")
	func1()

	result := calculator(10, 2)
	fmt.Println("Answer is", result)

	total := addition(10, 20, 6, 78)
	fmt.Println(total)

	ans, msg := multiReturns(10, 20)
	fmt.Printf("Answer is %v and message is %v", ans, msg)
}

func func1() {
	fmt.Println("Namasthe from go lang")
}

func calculator(num1 int, num2 int) int {
	return num1 + num2
}

//dont know how many value come. all value sum should calculate
func addition(params ...int) int {
	total := 0

	for _, value := range params {
		total += value // this equal to total = total + value
	}
	return total
}

//multiple returns
func multiReturns(val1 int, val2 int) (int, string) {
	return val1 + val2, "This will be the answer"
}

/*
	Rules
		1. function ekk tavat function ekk atule hadann ba
		2. anonymous function exists in go lang
*/
