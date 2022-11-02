package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Please enter your age")
	value, _ := reader.ReadString('\n')

	fmt.Println(value)

	coveretdToInt, err := strconv.ParseFloat(strings.TrimSpace(value), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added to 1 then now value is : ", coveretdToInt+1)
	}

}

/*
Explanation
	1. line 15 string type input ekk aran tiyenva. eke '\n' eken kiynne kochchr durata input ek gann onida kiyna eka.
	   mekata anuva new line ekkt yanakm
	2. 19 line eke krl tiyenne gann user input ek float ekkt convert krn ek. ehidi strings.Trimspace ek use krl tiyenne meka kelinm
	   convert krnn bari nisa. bcz mehi value ek tiyenne assume 4 input kra kiyla the 4\n lesa. so \n remove krnn tama ek use krl
	   tiyenne
	3. 21 krl tiyenne convertion eke error ek tiyeida bln ek. ek normal programing vala try catch vge. oni nm strings.TimeSpace
	   natuva run krl blnn etkota error ek print vevi

*/
