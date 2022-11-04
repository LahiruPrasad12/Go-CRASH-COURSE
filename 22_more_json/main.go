package main

import (
	"encoding/json"
	"fmt"
)

// Im not going to expoet it anywhere. thats why this start with simple letter
type course struct {
	Name     string `json:"coursename"` // me condition eken json eke name kiyn key ek venuvt course name vadinva
	Price    int
	PlatForm string   `json:"website"`        //meke Platform ekt website kiyla key ek vadinva
	Password string   `json:"-"`              // meken krl tiyenne password ek unselect krl
	Tags     []string `json:"tags,omitempty"` //meke name ek tags dala and null ek unselect krl
}

func main() {
	fmt.Println("Welcome to JSON in GoLang")

	encodeJSON()
}

func encodeJSON() {
	courses := []course{
		{"React JS Bootcamp", 1000, "www.react.com", "lahiru12", []string{"web", "react"}},
		{"Node JS Bootcamp", 1000, "www.node.com", "lahiru123", []string{"mobile", "node"}},
		{"Go Lang Bootcamp", 1000, "www.go.com", "lahiru1234", nil},
	}

	//pckage this data as JSON data
	finalJSON, err := json.Marshal(courses)                       //meka nikamm data penvai
	finalFormatJSON, err := json.MarshalIndent(courses, "", "\t") // meke data tika tabs space vlin format krl lassant penvanava
	checkError(err)

	fmt.Printf("%s\n", finalJSON)
	fmt.Printf("%s\n", finalFormatJSON)

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

/*
	me api hadapu JSON eke tiyana aul tama keys patan ganne capital letter vlin, password ek select venva. ek hide krnn oni and
	null data tiyn filed emat select venva. attatm nil tiyn ev select venn oni na ne. eca fixed krnn apita puluvn struct eke variable
	declare krn tanama conditions denn
*/
