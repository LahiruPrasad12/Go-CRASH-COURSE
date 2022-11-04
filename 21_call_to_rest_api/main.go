package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to call api using go lang")

	getData()
	createPostUSingJsonData()
	createPostUsingFormData()
}

// This is used to get post
func getData() {
	const URL = "http://localhost:3000/get"

	response, err := http.Get(URL)
	checkError(err)

	defer response.Body.Close()

	// fmt.Println(response)
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

// This is used to create post sending json type data
func createPostUSingJsonData() {
	const myurl = "http://localhost:3000/post"

	requestBody := strings.NewReader(`
		{
			"coursename":"ReactJS",
			"price" : 1000
		}
	`)

	response, err := http.Post(myurl, "application/json", requestBody)
	checkError(err)

	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)
	checkError(err)

	fmt.Println(string(content))

}

// This is used to create post sending formdata
func createPostUsingFormData() {
	const myurl = "http://localhost:3000/postform"

	data := url.Values{}
	data.Add("first_name", "lahiru")
	data.Add("last_name", "prasad")

	response, err := http.PostForm(myurl, data)
	checkError(err)

	defer response.Body.Close() // close the connection requests

	content, err := ioutil.ReadAll(response.Body)
	checkError(err)

	fmt.Println(string(content))

}

// This is used to handle errors in globaly
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
