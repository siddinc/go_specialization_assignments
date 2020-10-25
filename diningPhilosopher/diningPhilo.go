package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var wg sync.WaitGroup
var counter int = 0
var c1 chan bool = make(chan bool)

type Chopstick struct {
	sync.Mutex
}

type Philo struct {
	pairOfChopsticks [2]*Chopstick
	number           int
}

func (p *Philo) Eat() {
	go HostPermission(c1)
	j := <-c1

	if j {
		incrCounter()
		for i := 0; i < 3; i++ {
			x := rand.Intn(2)
			p.pairOfChopsticks[x].Lock()
			fmt.Println("starting to eat", p.number)
			p.pairOfChopsticks[x].Unlock()
			fmt.Println("finishing eating", p.number)
		}
		decrCounter()
	} else {
		fmt.Println("permission denied")
	}
	wg.Done()
}

func incrCounter() {
	counter++
}

func decrCounter() {
	counter--
}

func HostPermission(c chan bool) {
	if counter > 2 {
		c <- false
	}
	c <- true
}

func main() {
	chopsticks := make([]*Chopstick, 5)

	for i := 0; i < 5; i++ {
		chopsticks[i] = new(Chopstick)
	}

	philosophers := make([]*Philo, 5)

	for i := 0; i < 5; i++ {
		philosophers[i] = &Philo{[2]*Chopstick{chopsticks[i], chopsticks[(i+1)%5]}, i + 1}
	}

	wg.Add(5)
	for _, v := range philosophers {
		go v.Eat()
	}
	wg.Wait()
}
