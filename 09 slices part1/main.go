package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println("Welcome to slices")

	var fruitList = []string{"apple"}
	fmt.Println(fruitList) //-> [apple]

	//add data to slice
	fruitList = append(fruitList, "mango", "banana")
	fmt.Println(fruitList) //-> [apple Mango Banana]

	//get specific part of array
	speicific := append(fruitList[1:])
	fmt.Println(speicific) //->[Mango Banana]

	//get specific part of array (range selection)
	speicific = append(fruitList[1:2])
	fmt.Println(speicific) //->[Mango]

	speicific = append(fruitList[:2])
	fmt.Println(speicific) //->[apple Mango]

	//sorted string array
	sort.Strings(fruitList)
	fmt.Println("Sirted array is:", fruitList) //->[apple banana mango]

	//sorted int array
	var score = []int{}
	score = append(score, 2, 5, 3, 10, 4)
	fmt.Println("array is", score) //-> [2 5 3 10 4]
	sort.Ints(score)
	fmt.Println("Sorted array is:", score) // ->[2 3 4 5 10]

	//check whether the array is sorted or not
	fmt.Println(sort.StringsAreSorted(fruitList)) //-> true
	fmt.Println(sort.IntsAreSorted(score))        //->true

}

/*
	slice kiynne hriyata java vala arraylist ekk vge. meket size ek automatically expand venva
*/
