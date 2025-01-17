package concurrency_test

import (
	"sync"
	"testing"

	"github.com/AlejandroAldana99/challenge_homework/concurrency"
	"github.com/AlejandroAldana99/challenge_homework/stack"
)

func TestAddItemSafety(t *testing.T) {
	testStack := &stack.Stack{}
	var wg sync.WaitGroup

	// Items to be added concurrently
	items := []string{"item1", "item2", "item3", "item4", "item5"}

	// Add items concurrently
	for _, item := range items {
		wg.Add(1)
		go concurrency.AddItemSafety(item, testStack, &wg)
	}

	wg.Wait()

	// Validate that all items were added
	_, count := testStack.IsEmpty()
	if count != len(items) {
		t.Errorf("Expected stack size %d; got %d", len(items), count)
	}

	// Check the order of items in the stack
	expectedItems := make(map[string]bool)
	for _, item := range items {
		expectedItems[item] = true
	}

	for i := 0; i < len(items); i++ {
		top, ok := testStack.Pop()
		if !ok {
			t.Errorf("Expected to pop an item but got none")
		}

		if !expectedItems[top] {
			t.Errorf("Unexpected item in stack: %s", top)
		}
	}
}

func TestAddItemSafetyConcurrency(t *testing.T) {
	testStack := &stack.Stack{}
	var wg sync.WaitGroup

	// Test with a high number of concurrent goroutines
	totalItems := 100
	for i := 0; i < totalItems; i++ {
		wg.Add(1)
		go concurrency.AddItemSafety("item", testStack, &wg)
	}

	wg.Wait()

	// Validate items in the stack
	_, count := testStack.IsEmpty()
	if count != totalItems {
		t.Errorf("Expected stack size %d; got %d", totalItems, count)
	}
}
