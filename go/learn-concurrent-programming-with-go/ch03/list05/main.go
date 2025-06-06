package main

import (
	"fmt"
	"time"
)

func stingy(money *int) {
	for range 100000 {
		*money += 10
	}
	fmt.Println("Stingy Done")
}

func spendy(money *int) {
	for range 100000 {
		*money -= 10
	}
	fmt.Println("Spendy Done")
}

func main() {
	money := 100
	go stingy(&money)
	go spendy(&money)
	time.Sleep(2 * time.Second)
	fmt.Printf("Money in bank account: %d\n", money)
}
