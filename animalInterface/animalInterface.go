package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct{ name string }

func (c Cow) Eat() {
	fmt.Println("grass")
}

func (c Cow) Move() {
	fmt.Println("walk")
}

func (c Cow) Speak() {
	fmt.Println("moo")
}

type Snake struct{ name string }

func (s Snake) Eat() {
	fmt.Println("mice")
}

func (s Snake) Move() {
	fmt.Println("slither")
}

func (s Snake) Speak() {
	fmt.Println("hsss")
}

type Bird struct{ name string }

func (b Bird) Eat() {
	fmt.Println("worms")
}

func (b Bird) Move() {
	fmt.Println("fly")
}

func (b Bird) Speak() {
	fmt.Println("peep")
}

func main() {
	animalsMap := make(map[string]Animal)
back:
	fmt.Printf("Enter new request or X to exit > ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var Req string = scanner.Text()

	splitReq := strings.Split(Req, " ")

	for Req != "X" {
		switch splitReq[0] {
		case "newanimal":
			var a Animal

			switch splitReq[2] {
			case "cow":
				a = Cow{splitReq[1]}
			case "bird":
				a = Bird{splitReq[1]}
			case "snake":
				a = Snake{splitReq[1]}
			}

			animalsMap[splitReq[1]] = a
			fmt.Println("Created it!")

		case "query":
			i, ok := animalsMap[splitReq[1]]

			if !ok {
				fmt.Printf("%v not found. Please try again\n", splitReq[1])
			} else {
				switch splitReq[2] {
				case "eat":
					i.Eat()
				case "move":
					i.Move()
				case "speak":
					i.Speak()
				}
			}

		default:
			fmt.Println("Invalid Request. Please try again")
		}
		goto back
	}
}
