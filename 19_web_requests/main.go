package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {

	fmt.Println("Welcome to web requests")

	response, error := http.Get(url)
	errChecking(error)

	fmt.Println(response)

	defer response.Body.Close() //caller's responsibility to close the connection

	databyte, err := ioutil.ReadAll(response.Body)
	errChecking(err)
	fmt.Println(string(databyte))

}

func errChecking(err error) {
	if err != nil {
		fmt.Println("error")
	}
}
