package main

import (
	"fmt"
	"strings"
)

func main() {

	var ianString string

	fmt.Println("Please enter string: ")
	fmt.Scan(&ianString)

	if strings.HasPrefix(ianString, "i") && strings.Contains(ianString, "a") && strings.HasSuffix(ianString, "n") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
