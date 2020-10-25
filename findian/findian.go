package main

import (
	"fmt"
	"strings"
)

func findian(x string) int {
	y := strings.ToLower(x)

	if strings.HasPrefix(y, "i") && strings.HasSuffix(y, "n") && strings.Contains(y, "a") {
		return 1
	}
	return 0
}

func main() {
	var userStr string
	fmt.Printf("Enter a string: ")
	fmt.Scan(&userStr)

	if result := findian(userStr); result == 0 {
		fmt.Println("Not Found!")
	} else {
		fmt.Println("Found!")
	}
}
