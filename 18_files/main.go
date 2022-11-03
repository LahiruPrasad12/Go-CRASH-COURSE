package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("welcome to file")

	content := "This is first my file"

	file, err := os.Create("./my_file.txt")

	checkNilErr(err)

	length, err := io.WriteString(file, content)

	checkNilErr(err)

	fmt.Println("length is", length)
	defer file.Close()

	readFile("./my_file.txt")

}

// read file
func readFile(filename string) {
	dataByte, err := ioutil.ReadFile(filename)

	checkNilErr(err)

	fmt.Println(dataByte) //-> file ekk read krddi nitrm ek result ek tiyenne byte code ekk vidiyt so api ek string ekkt convert krnn oni
	fmt.Println(string(dataByte))
}

// check error globaly
func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
