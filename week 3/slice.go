package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var takeInp string

	fmt.Println("Please enter slice's int values seperated by ',' e.g. 7,4,9 :")
	fmt.Scan(&takeInp)

	newInput := strings.Split(takeInp, ",")

	toSortSlice := []int{}

	for _, val := range newInput {
		intVal, _ := strconv.Atoi(val)
		toSortSlice = append(toSortSlice, intVal)
	}

	for i := 0; i < len(toSortSlice)-1; i++ {
		minIndex := i
		for j := i; j < len(toSortSlice); j++ {
			if toSortSlice[j] < toSortSlice[minIndex] {
				minIndex = j
			}
		}
		temp := toSortSlice[minIndex]
		toSortSlice[minIndex] = toSortSlice[i]
		toSortSlice[i] = temp
	}
	fmt.Println(toSortSlice)

}
