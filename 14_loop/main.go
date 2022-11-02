package main

import "fmt"

func main() {

	fmt.Println("Welcome go loop")

	days := []string{"sunday", "monday", "tuesday", "wednesday"}

	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	//meka foreach with only index
	for i := range days {
		fmt.Println(days[i])
	}

	//foreach with value
	for index, value := range days {
		fmt.Printf("index is %v and value is %v \n", index, value)
	}

	value := 2

	for value < 10 {
		fmt.Println(value)
		value++
	}

	//break statement and continous statement
	newvalue := 1
	for newvalue < 10 {

		if newvalue == 2 {
			goto my_statement
		}

		if newvalue == 3 {
			newvalue++
			continue

		}

		if newvalue == 5 {
			break
		}
		fmt.Println(newvalue)
		newvalue++
	}

	//statement
my_statement:
	{
		fmt.Println("jumping at 2")
		fmt.Println("jumping at 3")
	}

}

/*
	i++ puluvn unata anit languages vala vge meke ++i support krnne na;
	metana api create krpu statement ek hriyata func tion ekk vge. apita mema kamati name ekk dila atule mkkhri kriyavaliyak krnn puluvn
*/
