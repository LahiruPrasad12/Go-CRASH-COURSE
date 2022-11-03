package main

import "fmt"

func main() {

	defer fmt.Println("go lang")
	fmt.Println("welcome to deffer")
	/*->
	out put will be
		welcome to deffer
		go lang
	*/

	defer fmt.Println("one")
	defer fmt.Println("two")
	fmt.Println("three")
	/*->
	out put will be
		welcome to deffer
		three
		two
		one
		go lang
	*/

	myDiffer()
	/*->
	out put will be
		welcome to deffer
		4321
		two
		one
		go lang
	*/
}

func myDiffer() {
	for num := 1; num < 5; num++ {
		defer fmt.Print(num)
	}
}

/*
	deffer eke saralavam venne meka. program ek run venne line by line ne. ehidi apita mulin liuva oni code ekk avasanat run krnn oni nm
	me differ use krnn puluvn. e qwe differe dapu line ekk nitrm run venne program ek avasanetm.

	differ ek use krnnet last in first out method ek. stack method ek tama differ use krnnet

	sralavama differ krnne differ dapu code stack ekkt dala program ek avasanetm run krn eka
*/
