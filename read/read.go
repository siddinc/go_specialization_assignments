package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Name struct {
	fname string
	lname string
}

var FileName string

func main() {
	fmt.Printf("Enter the filename: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	FileName = scanner.Text()

	names, err := readFile(FileName)

	if err != nil {
		fmt.Println(err)
	}

	for _, v := range names {
		fmt.Printf("Fname: %s\tLName: %s\n", v.fname, v.lname)
	}
}

func readFile(fileName string) ([]Name, error) {
	file, err := os.Open(fileName)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	names := make([]Name, 0, 2)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		splitName := strings.Split(scanner.Text(), " ")

		if len(splitName[0]) > 20 || len(splitName[1]) > 20 {
			return nil, errors.New("Name not under 20 characters")
		}

		names = append(names, Name{fname: splitName[0], lname: splitName[1]})
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return names, nil
}
