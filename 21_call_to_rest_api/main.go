package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to call api using go lang")
	getData()
}

func getData() {
	const URL = "http://localhost:3000/get"

	response, err := http.Get(URL)
	checkError(err)

	defer response.Body.Close()

	fmt.Println(response)
	fmt.Println("Status code ", response.StatusCode)      // Status code  200
	fmt.Println("Content length", response.ContentLength) //Content length 34

	content, err := ioutil.ReadAll(response.Body)

	checkError(err)

	//byte code conver to string part 1
	fmt.Println(string(content)) //{"message":"Hello from learncode"}

	//byte code conver to string part 2
	var responseString strings.Builder
	byteCount, _ := responseString.Write(content)

	fmt.Println("Byte count is", byteCount) //Byte count is 34
	fmt.Println(responseString.String())    //{"message":"Hello from learncode"}

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
