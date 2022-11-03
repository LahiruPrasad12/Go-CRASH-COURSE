package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://localhost:3000/learn?coursename=reactjs"

func main() {

	fmt.Println("welcome to urls in go langs")

	//parsing
	result, _ := url.Parse(myurl) //me url ek nikm tiyenne string ekk vge. so api ek url ekkt conver krgnn oni

	//mema url ek convert krgttama eken apita avashaya oni value ekk gann puluvn. As example
	fmt.Println(result)          //-> https://localhost:3000/learn?coursename=reactjs
	fmt.Println(result.Scheme)   //-> https
	fmt.Println(result.Host)     //-> localhost:3000
	fmt.Println(result.Path)     //-> /learn
	fmt.Println(result.RawQuery) //-> coursename=reactjs
	fmt.Println(result.Port())   //-> 3000

	//get params values
	quaryparams := result.Query()
	fmt.Println(quaryparams["coursename"]) //[reactjs]

	//foreach values
	for _, val := range quaryparams {
		fmt.Println(val)
	}

	//api me venkm baluve URL ekk data access krn ek. metanin passe blmu avashya data tiyeddi komada ekt URL ekk hadanne kiyla
	parseURL := &url.URL{
		Scheme:  "https",
		Host:    "localhost:3000",
		Path:    "/learn",
		RawPath: "user=lahiru",
	}

	newURL := parseURL.String()
	fmt.Println(newURL)

}
