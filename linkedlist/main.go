package main

import (
	"fmt"
	"github.com/obihann/GoDataStructures/linkedlist/src"
)

func main() {
	list := LinkedList.NewList()
	size := 20

	for x := 0; x < size; x++ {
		if err := list.Append(x); err != nil {
			fmt.Printf("An error occured while trying to append to the list: %s\n", err)
		} else {
			fmt.Printf("adding %v\n", (x + 1))
		}
	}

	if err := list.Delete(3); err != nil {
		fmt.Printf("An error occured while attempting to delete %v: %s\n", 3, err)
	}

	if err := list.Delete(9); err != nil {
		fmt.Printf("An error occured while attempting to delete %v: %s\n", 9, err)
	}

	if err := list.Delete(4); err != nil {
		fmt.Printf("An error occured while attempting to delete %v: %s\n", 4, err)
	}

	if err := list.Delete(7); err != nil {
		fmt.Printf("An error occured while attempting to delete %v: %s\n", 7, err)
	}

	arr := list.List()

	for x := 0; x < len(arr); x++ {
		fmt.Printf("%v\n", (arr[x] + 1))
	}
}
