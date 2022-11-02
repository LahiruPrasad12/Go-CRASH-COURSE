package main

import "fmt"

func main() {

	var ptr1 *int
	fmt.Println("Value of pointer is : ", ptr1) //-> this will print nil. bcz still doesnt assign any value

	myNumber := 25
	var ptr2 = &myNumber
	fmt.Println("Memory address of myNumber variable : ", ptr2)  //-> This will print the memory address of variable
	fmt.Println("Memory address of myNumber variable : ", *ptr2) //-> This will print the value of variable

	*ptr2 = *ptr2 + 2
	fmt.Println("New value is : ", myNumber) //-> This will print the 27

}
