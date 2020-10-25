package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	var name, address string

	fmt.Printf("Enter the name: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	name = scanner.Text()

	fmt.Printf("Enter the address: ")
	scanner = bufio.NewScanner(os.Stdin)
	scanner.Scan()
	address = scanner.Text()

	m := map[string]string{
		"name":    name,
		"address": address,
	}

	var b []byte
	b, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(b))
}
