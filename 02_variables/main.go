package main

import "fmt"

func main() {

	/*variable declaration*/
	var name string = "Lahiru"
	fmt.Println(name)
	fmt.Printf("Variable is type of : %T \n", name)

	/*implicity declaration*/
	var age = "Lahiru"
	isApproved := true
	fmt.Println(isApproved, age) // -> both are correct as implicity declaration
	fmt.Printf("Variable type is of : %T ", isApproved)

	/*Default values*/
	var bol bool // -> default value ek false ve, String vala "" and int vala 0 ve
	fmt.Println(bol)

	/*constant varaible*/
	const JwtToken string = "7nnee64ghb"
	fmt.Println(JwtToken)
	fmt.Printf("Variable type is of : %T \n", JwtToken)

}

/*
Keep In Mind These things
	implicity declartion method ekk atule vitri apita mema variable declare krnn puluvn. method ekkin pita gloably apita melesa
	variable decalre krnn ba. ehidi apita var keyword ek use krnnm venva

	variable name ekk capital letter ekkin patan gannavnm ek public variable ekk venva. java vala public keyword dala varaible hdnva
	vge tama GoLang eke capital letter ekkin variable ekk hadan ek. Then me "JwtToken" variable ek apita ona go file ekk idn access
	krnn puluvn
*/
