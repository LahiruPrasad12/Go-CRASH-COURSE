package main

import "fmt"

func main() {
	fmt.Println("welcome to structs")

	//This is like objcet
	me := User{"Lahiru", "lahirup471@gmail.com", true, 16}

	fmt.Println(me)      //-> {Lahiru lahirup471@gmail.com true 16}
	fmt.Println(me.Name) //-> Lahiru

	fmt.Printf("%v\n", me)  //-> {Lahiru lahirup471@gmail.com true 16}
	fmt.Printf("%+v\n", me) //-> {Name:Lahiru Email:lahirup471@gmail.com Status:true Age:16}. simply + mark print more details abt variable

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

/*
	1. structs kiynne java vala tiyna classes vge
	2. User eke variable hama ekkm capital letter ekkin patan gena tiyenne. its means eva public. oni tanak idn eva access krnn puluvn
	3. In No inheritance in go lang. No super or parent
	4.
*/
