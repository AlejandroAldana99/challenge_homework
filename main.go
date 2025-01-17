// Implement and Use a Stack in Go
package main

import (
	"fmt"
	"sync"

	"github.com/AlejandroAldana99/challenge_homework/concurrency"
	"github.com/AlejandroAldana99/challenge_homework/stack"
)

func main() {
	// Basic approach
	myStack := stack.Stack{}

	myStack.Push("Hello")
	isEmpty, nItems := myStack.IsEmpty()
	fmt.Printf("Is empty: %t, number of items: %d \n", isEmpty, nItems)

	item, found := myStack.Peek()
	fmt.Printf("Iteam found: %s, was found: %t \n", item, found)

	item, found = myStack.Pop()
	fmt.Printf("Item removed: %s, was found: %t \n", item, found)

	isEmpty, nItems = myStack.IsEmpty()
	fmt.Printf("Final is empty: %t, number of items: %d \n\n", isEmpty, nItems)

	// Add concurrency
	fmt.Println("Adding concurrency")
	var wg sync.WaitGroup

	for i := 0; i <= 3; i++ {
		wg.Add(1)
		value := fmt.Sprintf("test_%d", i)
		go concurrency.AddItemSafety(value, &myStack, &wg)
	}

	wg.Wait()

	isEmpty, nItems = myStack.IsEmpty()
	fmt.Printf("Concurrency results -> is empty: %t, number of items: %d \n", isEmpty, nItems)
}
