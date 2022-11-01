package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("welcome to time study of go lang")

	presentTime := time.Now()

	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006")) //me time ek me vidiytm denn oni. current date ek print kri
	fmt.Println(presentTime.Format("15:04:05"))   //me time ek me vidiytm denn oni. current time ek print kri
	fmt.Println(presentTime.Format("Monday"))     //me time ek me vidiytm denn oni. current day ek print kri

	createdDate := time.Date(2020, time.April, 10, 23, 23, 0, 0, time.UTC)
	fmt.Println(createdDate) /*This will print 2020-04-10 23:23:00 +0000 UTC. previous date ekk ganne memai. mehidi api dena date ek print venne. mekt oni nm apita format krnn puluvn
	kamati vidiykt. but ekttat ara pai dunn format time datem denn oni. like this,*/
	fmt.Println(createdDate.Format("01-02-2006 Monday")) // This will print 04-10-2020 Friday

}

/*
Date time format ekedi sample type ekt me dila tiyna date time day m denn oni. otherwise current date time enne na. varadi ev enne
*/
