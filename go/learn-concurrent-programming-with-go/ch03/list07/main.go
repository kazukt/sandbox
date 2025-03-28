package main

import (
	"fmt"
	"runtime"
	"time"
)

func stingy(money *int) {
	for range 100000 {
		*money += 10
		runtime.Gosched()
	}
	fmt.Println("Stingy Done")
}

func spendy(money *int) {
	for range 100000 {
		*money -= 10
		runtime.Gosched()
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
