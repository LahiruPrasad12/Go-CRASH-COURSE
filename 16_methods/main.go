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
func (u *User) getName() {
	fmt.Println("Hi", u.Name)
}

//create function to change the user name
func (u *User) newName() {
	u.Name = "bandara"
	fmt.Println("New name is ", u.Name)
}

func test(name *int) {
	*name = 10
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

/*
	uda eke aul kihipayak ava.

	1. newName() eke pointer variabvle ekk haduvata *u.Name use kre na value ek assign krnn. ekt hetuva You don't need to use *u.Name = "bandara" because u.Name is syntactic sugar in Go
	for (*u).Name.
	attama newName() kiynne method ekk etkota method ekk akara dekai tiyenne. value reciver and ponter reciver.

		1. value reciver method: 	func (u User) newName() {
										u.Name = "bandara"
										fmt.Println("New name is ", u.Name)
									}

		2. Pointer reciver method: 	func (u *User) newName() {
										u.Name = "bandara"
										fmt.Println("New name is ", u.Name)
									}


		meke value reciver ek compile and run una namut originam struct ek change krnne na. bcz metana u kiynne copy of the User struct. Modifying u.Name changes the copy,
		not the original struct.

		pointer reciever valadi u kiynne User struct eke copy ekk nevei. So ek Directly modifies the original User struct because it works with a pointer.

		etkota pointer receiver ekkdi api *u.Name = "Bandara" dann oni na. bcz pointer receiver ekkdi Go la apita ek automatically krl denva.


	Etkota meke test kiynne method ekk nevei function ekk. so eke recivers na. eke krl tiyenne pointer type argument ekk aran name kiyl. itin api ekt value ekk use krnvnm apit
	*name = 10 vge denn venva

*/
