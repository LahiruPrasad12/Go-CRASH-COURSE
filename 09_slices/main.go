package main

import "fmt"

func main() {

	fmt.Println("Welcome to slices")

	var fruitList = []string{"apple"}
	fmt.Println(fruitList) //-> [apple]

	//add data to slice
	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList) //-> [apple Mango Banana]

	//get specific part of array
	speicific := append(fruitList[1:])
	fmt.Println(speicific) //->[Mango Banana]

	//get specific part of array (range selection)
	speicific = append(fruitList[1:2])
	fmt.Println(speicific) //->[Mango]

	speicific = append(fruitList[:2])
	fmt.Println(speicific) //->[apple Mango]
}

/*
	slice kiynne hriyata java vala arraylist ekk vge. meket size ek automatically expand venva
*/
