package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Println("Welcome to switch case")

	rand.Seed(time.Now().UnixNano()) // me line ek nattan nitrm value ek vadinne 6. meka damma 0-6 enkm value randomly vadinva
	diceNumber := rand.Intn(6) + 1   // +1 damme nattan meke random value vadinne 0-5 venkm
	fmt.Println(diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Value is 1")
	case 2:
		fmt.Println("Value is 2")
		fallthrough
	default:
		fmt.Println("Value is greater")
	}

}

/*
	1. Meke udama seeder ekk dala tiyenne apita random value gann nikm adarakayak vage
	2. fallthrough eken krnne me case eke fall unat program ek terminate krnne natuva digatam run rkn eka
		EX:- random value ek vidiyt 2 avot "Value is 2" and "Value is greater" ekt print venva.
			 1 vadunot "Value is 1" and "Value is 2" lesa out put ek vadinva
*/
