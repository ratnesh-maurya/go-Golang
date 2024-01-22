package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("welcome")

	presenttime := time.Now()

	fmt.Println("It's currently ", presenttime)

	fmt.Println(presenttime.Format("01-02-2006 15:04:05 Monday"))
	fmt.Println(presenttime.Format("01-02-2006 Monday"))

	createddate := time.Date(2020, time.August, 10, 23, 23, 0, 0, time.UTC)

	fmt.Println(createddate)

	fmt.Println(createddate.Format("01-02-2006 Monday"))

	
}
