package main

import (
	"fmt"
	"time"
)

func goRoutine1(x *int) {
	*x = 4
}

func goRoutine2(x *int) {
	*x = 6
}

func main() {
	var x int = 1

	go goRoutine1(&x)
	go goRoutine2(&x)

	time.Sleep(10 * time.Millisecond)
	fmt.Println(x)
}

/* Here goRoutine1 and goRoutine2 are in race condition. The initial value of int x = 1.
When we pass x by reference to goRoutine1, it modifies the value of x to 4.
When we pass x by reference to goRoutine2, it modifies the value of x to 6.
Hence the latest updated value of x will depend on the order of execution of goRoutine1 and goRoutine2.*/
