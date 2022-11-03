package main

import "fmt"

func main() {

	var courses = []string{"react-js", "javascript", "swift", "python"}
	fmt.Println(courses)

	//remove value from slice
	index := 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
