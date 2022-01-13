package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	names := [][]string{{"fname1", "lname1"}, {"fname2", "lname2"}} //name slice

	fmt.Printf("Enter file name with extension eg. filename.txt : ")
	var fileName string
	fmt.Scan(&fileName)

	file, _ := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644) //if file exist append else create new file

	for i := 0; i < len(names); i++ {
		for j := 0; j < len(names); j++ {
			file.WriteString(names[i][j] + " ")
		}
		file.WriteString("\n")
	}

	byt, _ := ioutil.ReadFile(fileName) //reading the file
	defer file.Close()

	fmt.Println(string(byt)) //printing the file content
}
