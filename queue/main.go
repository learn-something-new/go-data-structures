package main

import (
	"fmt"
	"github.com/obihann/GoDataStructures/queue/src"
)

func main() {
	queue := Queue.NewQueue()
	size := 15

	for x := 0; x < size; x++ {
		fmt.Printf("Adding %v\n", x)
		queue.Add(x)
	}

	for x := 0; x < size; x++ {
		num := queue.Del()

		if num == -1 {
			fmt.Printf("Del failed! Queue is empty.")
		} else {
			fmt.Printf("Deleted %v\n", num)
		}
	}
}
