package main

import (
	"fmt"
	"github.com/obihann/GoDataStructures/linkedlist/src"
)

func main() {
	list := LinkedList.NewList()

	for x := 0; x <= 10; x++ {
		fmt.Printf("adding %v\n", x)
		list.Append(x)
	}

	list.Delete(3)
	list.Delete(9)
	list.Delete(4)

	arr := list.List()

	for x := 0; x < len(arr); x++ {
		fmt.Printf("%v\n", arr[x])
	}
}
