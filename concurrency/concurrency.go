package concurrency

import (
	"fmt"
	"sync"

	"github.com/AlejandroAldana99/challenge_homework/stack"
)

var mutex sync.Mutex

func AddItemSafety(value string, stack *stack.Stack, wg *sync.WaitGroup) {
	defer wg.Done()

	mutex.Lock()
	fmt.Printf("Pushing element %s to Stack \n", value)
	stack.Push(value)
	fmt.Printf("Element %s pushed to Stack \n", value)
	mutex.Unlock()
}
