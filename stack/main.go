package main

import (
	"fmt"
	"github.com/obihann/GoDataStructures/stack/src"
)

func main() {
	stack := Stack.NewStack()
	size := 10

	for x := 0; x < size; x++ {
		fmt.Printf("Pushing %v\n", x)
		success := stack.Push(x)

		if !success {
			fmt.Printf("Push failed! Stack is full.")
		}
	}

	for x := 0; x < size; x++ {
		num := stack.Pop()

		if num == -1 {
			fmt.Printf("Pop failed! Stack is empty.")
		} else {
			fmt.Printf("Popped %v\n", num)
		}
	}
}
