package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Rate the service")

	// Comma ok symble(err ok). mema user input gann ekt api kiynne comma ok or err ok sysmbel ek kiyla

	input, _ := reader.ReadString('\n')
	fmt.Println(input)

}

/*
Keep mind thes things
	Go language vala try catch na. api error ekk avot ek variable ekkt assign krgnne. assume user input ek gann eke error ekk avot apita
	_ lakuna venuvt variable name ekk dila e error ek e variable ekt assign krgnn puluvn
*/
