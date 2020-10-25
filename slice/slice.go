package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	s := make([]int, 0, 3)
	var i string

back:
	fmt.Printf("Enter an int: ")
	fmt.Scan(&i)

	for i != string('X') {
		j, _ := strconv.Atoi(i)
		s = append(s, j)
		sort.Ints(s)
		fmt.Println("The sorted slice:", s)
		goto back
	}

	fmt.Println("Exiting")
}
