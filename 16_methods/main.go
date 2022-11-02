package main

import "fmt"

func main() {
	fmt.Println("welcome to go lang methods")

	user := User{"Lahiru", 20}
	fmt.Println(user)

	//call to method
	user.getName()

	user.newName()
	fmt.Println(user)
}

type User struct {
	Name string
	Age  int
}

//create a methods
func (u User) getName() {
	fmt.Println("Hi", u.Name)
}

//create function to change the user name
func (u User) newName() {
	u.Name = "bandara"
	fmt.Println("New name is ", u.Name)
}

/*
	Methods and function vala vensk na. function ekk class ekk atule define krm api ek method ekk vidiyat hadunvanne

	regarding last function newName. that used to change property.
		attatm meka run krm enne me vge answer ekk
			{Lahiru 20}, New name is  bandara, {Lahiru 20}
		meka baluvama penva api name ek vens krnn kalin object eke tiyenne lahiru name eke value ek vidiyt nut change krt passe tiyennet
		lahiru. meyt hetuvt tama mehi value ekk copy ekk pass kirima. itin menn mekt solution ekk vidiyt tama pointers enne. apita attam
		real value ek pass krnn oni nm pointers tama use krnn oni. otherwise samahara velavat me una vge real value ek natuva
		value eke copyk pass venva
*/
