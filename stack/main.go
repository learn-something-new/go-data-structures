package main

import (
	"fmt"
	"github.com/obihann/GoDataStructures/stack/src"
)

func main() {
	stack := Stack.NewStack()
	size := 20

	for x := 0; x < size; x++ {
		fmt.Printf("Pushing %v\n", (x + 1))

		stack.Push(x + 1)
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
