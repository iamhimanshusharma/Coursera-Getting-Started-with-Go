package main

import "fmt"

func main() {

	var floatNum float32

	fmt.Println("Please enter floating point number: ")
	fmt.Scan(&floatNum)

	fmt.Println("Truncated Integer number is: ", int32(floatNum))
}
