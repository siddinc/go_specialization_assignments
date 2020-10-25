package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func sortBatch(s []int, c chan []int) {
	sort.Ints(s)
	c <- s
}

func splitIntoBatches(s []int, batchSize int) [][]int {
	var chunk []int
	chunks := make([][]int, 0)

	for len(s) >= batchSize {
		chunk, s = s[:batchSize], s[batchSize:]
		chunks = append(chunks, chunk)
	}

	if len(s) > 0 {
		chunks = append(chunks, s[:len(s)])
	}

	return chunks
}

func main() {
	var sequence string
	c := make(chan []int)
	intSequence := make([]int, 0)

	fmt.Printf("Enter a sequence of 12 integers separated by spaces: ")
	reader := bufio.NewReader(os.Stdin)
	sequence, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
	}

	sequence = strings.Replace(sequence, "\n", "", -1)
	splitSequence := strings.Split(sequence, " ")

	for _, v := range splitSequence {
		i, err := strconv.Atoi(v)

		if err != nil {
			return
		}

		intSequence = append(intSequence, i)
	}

	batches := splitIntoBatches(intSequence, 3)
	n := make([]int, 0)

	for i := 0; i < len(batches); i++ {
		go sortBatch(batches[i], c)
		x := <-c
		fmt.Printf("goroutine %v result is: %v\n", i+1, x)
		n = append(n, x...)
	}

	sort.Ints(n)
	fmt.Println("The final sorted array is:", n)
}
