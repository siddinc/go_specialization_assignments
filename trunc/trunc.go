package main

import "fmt"

func truncate(x float64) int {
	var result int = int(x)
	return result
}

func main() {
	var (
		fpNum    float64
		truncNum int
	)

	fmt.Printf("Enter a floating point no.: ")
	fmt.Scan(&fpNum)

	truncNum = truncate(fpNum)
	fmt.Println("The truncated no. is", truncNum)
}
