package main

import (
	"fmt"
	"time"
)

func main() {

	n := time.Now()
	fmt.Println("I am printing this at ", n)

	t := time.Date(2022, time.February, 15, 10, 50, 0, 0, time.UTC)
	fmt.Println("Go launched at ", t)
	fmt.Println(t.Format(time.ANSIC))

}
