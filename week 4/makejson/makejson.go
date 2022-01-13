package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

func main() {
	var typeName string
	var typeAddress string
	var p1 Person

	fmt.Printf("name:")
	fmt.Scan(&typeName)
	fmt.Printf("address:")
	fmt.Scan(&typeAddress)

	p1 = Person{Name: typeName, Address: typeAddress}

	byte, _ := json.Marshal(p1)

	fmt.Printf(string(byte))
}
