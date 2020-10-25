package main

import (
	"fmt"
)

var lenSlice int

func main() {
	var x int
back:
	fmt.Printf("Enter length of sequence upto 10: ")
	fmt.Scan(&lenSlice)

	if lenSlice > 10 {
		fmt.Println("Length of sequence cannot be > 10. Try again")
		goto back
	}

	numbers := make([]int, 0, lenSlice)

	for i := 0; i < lenSlice; i++ {
		fmt.Printf("Enter integer %v: ", i+1)
		fmt.Scan(&x)
		numbers = append(numbers, x)
	}
	fmt.Println("The unsorted sequence is:", numbers)

	bubbleSort(numbers)
	fmt.Printf("The sorted sequence is: ")

	for i := 0; i < lenSlice; i++ {
		fmt.Printf("%d ", numbers[i])
	}
}

func bubbleSort(numbers []int) {
	for i := 0; i < lenSlice; i++ {
		for j := 0; j < lenSlice-i-1; j++ {
			swap(numbers, j)
		}
	}
}

func swap(numbers []int, i int) {
	if numbers[i] > numbers[i+1] {

		temp := numbers[i]
		numbers[i] = numbers[i+1]
		numbers[i+1] = temp
	}
}
