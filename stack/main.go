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

		fmt.Printf("The count is %v\n", stack.Top())
		fmt.Printf("The cap is %v\n", stack.Len())
		fmt.Printf("The length is %v\n\n", stack.Cap())
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
