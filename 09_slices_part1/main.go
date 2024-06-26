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
	keep this mind. nitaram [:3] vge tibbot eliment 3k select krl 3 veni eka ain krnva. ek tama behind scenario eka

	speicific = append(fruitList[:2]) saha speicific = fruitList[:2] menn me deke venasa tama api pahala kata krl tiyenne. meke saralavam venne append keyword ek damma specific slice ek
	hadenne new slice ekk vidiyt (speicific is a new slice that contains a copy of the first two elements of fruitList).

	but append keyword ek natuva tiyenne specific slice ek fruitlist reference vela tama hadenne. etkota api specific eke mokk hri change ekk krot ek fruitList ektat affect venva
	(Here, speicific is a slice that references the first two elements of fruitList. speicific and fruitList share the same underlying array.)



	Let's assume fruitList is defined as follows:


	fruitList := []string{"apple", "banana", "cherry", "date"}
	Example 1: Without append

	speicific = fruitList[:2]
	fmt.Println(speicific) // Output: [apple banana]
	Here, speicific is a slice that references the first two elements of fruitList.
	speicific and fruitList share the same underlying array.
	Now, let's modify an element in speicific:


	speicific[0] = "apricot"
	fmt.Println(speicific) // Output: [apricot banana]
	fmt.Println(fruitList) // Output: [apricot banana cherry date]
	Changing speicific[0] also changes fruitList[0] because they share the same underlying array.


	Example 2: With append

	speicific = append(fruitList[:2])
	fmt.Println(speicific) // Output: [apple banana]
	Here, speicific is a new slice that contains a copy of the first two elements of fruitList.
	speicific has its own underlying array separate from fruitList.
	Now, let's modify an element in speicific:

	speicific[0] = "apricot"
	fmt.Println(speicific) // Output: [apricot banana]
	fmt.Println(fruitList) // Output: [apple banana cherry date]
	Changing speicific[0] does not affect fruitList because speicific has its own underlying array.
*/
