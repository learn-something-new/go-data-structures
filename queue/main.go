package main

import (
	"fmt"
	"github.com/obihann/GoDataStructures/queue/src"
)

func main() {
	queue := Queue.NewQueue()
	size := 20

	for x := 0; x < size; x++ {
		if err := queue.Add(x + 1); err != nil {
			fmt.Printf("An error occured while trying to push to the stack: %s\n", err)
			break
		} else {
			fmt.Printf("Adding %v\n", (x + 1))
		}
	}

	for x := 0; x < size; x++ {
		num, err := queue.Del()

		if err != nil {
			fmt.Printf("Del failed! Queue is empty.")
		} else {
			fmt.Printf("Deleted %v\n", (num + 1))
		}
	}
}
