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
	isEmpty, count := myStack.IsEmpty()
	fmt.Printf("Is empty: %t, number of items: %d \n", isEmpty, count)

	item, found := myStack.Peek()
	fmt.Printf("Iteam found: %s, was found: %t \n", item, found)

	item, found = myStack.Pop()
	fmt.Printf("Item removed: %s, was found: %t \n", item, found)

	isEmpty, count = myStack.IsEmpty()
	fmt.Printf("Final is empty: %t, number of items: %d \n\n", isEmpty, count)

	// Add concurrency
	fmt.Println("Adding concurrency")
	var wg sync.WaitGroup
	// items to add concurrently
	items := []string{"item1", "item2", "item3", "item4", "item5"}

	for _, item := range items {
		wg.Add(1)
		go concurrency.AddItemSafety(item, &myStack, &wg)
	}

	wg.Wait()

	// Check items in the stack
	isEmpty, count = myStack.IsEmpty()
	fmt.Printf("Concurrency results -> is empty: %t, number of items: %d \n", isEmpty, count)
}
