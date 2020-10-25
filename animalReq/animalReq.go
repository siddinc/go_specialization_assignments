package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

var AnimalMap map[string]Animal = map[string]Animal{
	"cow":   {"grass", "walk", "moo"},
	"bird":  {"worms", "fly", "peep"},
	"snake": {"mice", "slither", "hsss"},
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.noise)
}

var Req string

func main() {
back:
	fmt.Printf("Enter new request or X to exit > ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	Req = scanner.Text()

	for Req != "X" {
		splitReq := strings.Split(Req, " ")
		animal, info := splitReq[0], splitReq[1]

		x := AnimalMap[animal]

		switch info {
		case "eat":
			x.Eat()
		case "move":
			x.Move()
		case "speak":
			x.Speak()
		}
		goto back
	}
}
