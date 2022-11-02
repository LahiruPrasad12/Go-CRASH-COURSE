package main

import "fmt"

func main() {

	fmt.Println("Wlcome to array in go lang")

	//just declare variable
	var fruitList [4]string

	fruitList[0] = "apple"
	fruitList[1] = "mango"
	fruitList[3] = "tomato"

	fmt.Println("Array is ", fruitList)             //-> This will print Array is [apple mango  tomato]. 2nd element ekt space ekk vadila
	fmt.Println("Array length is ", len(fruitList)) //-> This will print 4. attam array eke value 3k tibbad length ek 4 lesa vadinne

	//declare and define variable
	var vegList = [3]string{"potato", "beans"}
	fmt.Println("Array is ", vegList)

}
