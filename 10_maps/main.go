package main

import "fmt"

func main() {

	fmt.Println("welcome to maps")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println(languages) //-> map[JS:Javascript PY:Python RB:Ruby]

	//select one value
	fmt.Println(languages["JS"]) //-> Javascript

	//delete values
	delete(languages, "JS")
	fmt.Println(languages) //-> map[PY:Python RB:Ruby]

	//iterate map
	for key, value := range languages {
		fmt.Printf("Key is %v & Value is %v\n", key, value)
	}
	/*  Key is RB & Value is Ruby
	    Key is PY & Value is Python
	*/

}

/*
Keep These in mind
	1. map() kiynne key value pairs vlin data save krgnn ekt. java vala map ek vgemai
	2. make() ek atule tama map ek hadal tiyenne. nattan apita new() ek atulet meka hdann puluvn
	3. map(string)string kiynne key ek string and value ekt string type kiyn eka. varahan atule denne key type ek. pita denne value
	   type ek
	4. %v means value itself. apita value ekk print krgnn oni nm %v d value ekk type ek oni nm %T d use krnn puluvn

*/
