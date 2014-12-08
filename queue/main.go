package main

import (
	"fmt"
	"github.com/obihann/GoDataStructures/queue/src"
)

func main() {
	queue := Queue.NewQueue()
	size := 20

	for x := 0; x < size; x++ {
		fmt.Printf("Adding %v\n", (x + 1))

		queue.Add(x + 1)

		fmt.Printf("End is %v\n", queue.End())
		fmt.Printf("Len is %v\n", queue.Len())
		fmt.Printf("Cap %v\n\n", queue.Cap())
	}

	fmt.Printf("Deleted %v\n", queue.Del())

/*
 *    for x := 0; x < size; x++ {
 *        num := queue.Del()
 *
 *        if num == -1 {
 *            fmt.Printf("Del failed! Queue is empty.")
 *        } else {
 *            fmt.Printf("Deleted %v\n", num)
 *        }
 */
	}
}
