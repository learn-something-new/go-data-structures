package main

import (
	"fmt"
	"github.com/obihann/GoDataStructures/stack/src"
)

func main() {
	var stack Stack.Stack

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	stack.Push(6)
	stack.Push(7)
	stack.Push(8)
	stack.Push(9)
	stack.Push(10)

	fmt.Printf("%v\n", stack.Pop())
}
