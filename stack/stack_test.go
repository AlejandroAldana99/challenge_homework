package stack_test

import (
	"testing"

	"github.com/AlejandroAldana99/challenge_homework/stack"
)

func TestIsEmpty(t *testing.T) {
	s := stack.Stack{}

	// Test when stack is empty
	empty, count := s.IsEmpty()
	if !empty || count != 0 {
		t.Errorf("Expected IsEmpty to return true, 0; got %v, %d", empty, count)
	}

	// Test when stack is not empty
	s.Push("item1")
	empty, count = s.IsEmpty()
	if empty || count != 1 {
		t.Errorf("Expected IsEmpty to return false, 1; got %v, %d", empty, count)
	}
}

func TestPush(t *testing.T) {
	s := stack.Stack{}

	// Test adding items
	s.Push("item1")
	if _, count := s.IsEmpty(); count != 1 {
		t.Errorf("Expected stack size 1; got %d", count)
	}

	s.Push("item2")
	if _, count := s.IsEmpty(); count != 2 {
		t.Errorf("Expected stack size 2; got %d", count)
	}
}

func TestPop(t *testing.T) {
	s := stack.Stack{}

	// Test popping from an empty stack
	item, ok := s.Pop()
	if ok || item != "" {
		t.Errorf("Expected Pop to return \"\", false; got %q, %v", item, ok)
	}

	// Test popping items from the stack
	s.Push("item1")
	s.Push("item2")

	item, ok = s.Pop()
	if !ok || item != "item2" {
		t.Errorf("Expected Pop to return \"item2\", true; got %q, %v", item, ok)
	}

	item, ok = s.Pop()
	if !ok || item != "item1" {
		t.Errorf("Expected Pop to return \"item1\", true; got %q, %v", item, ok)
	}

	// Ensure the stack is empty after popping all items
	if empty, count := s.IsEmpty(); !empty || count != 0 {
		t.Errorf("Expected stack to be empty after popping all items; got %v, %d", empty, count)
	}
}

func TestPeek(t *testing.T) {
	s := stack.Stack{}

	// Test peeking on an empty stack
	item, ok := s.Peek()
	if ok || item != "" {
		t.Errorf("Expected Peek to return \"\", false; got %q, %v", item, ok)
	}

	// Test peeking without removing the item
	s.Push("item1")
	s.Push("item2")

	item, ok = s.Peek()
	if !ok || item != "item2" {
		t.Errorf("Expected Peek to return \"item2\", true; got %q, %v", item, ok)
	}

	// Ensure Peek does not remove the item
	if _, count := s.IsEmpty(); count != 2 {
		t.Errorf("Expected stack size 2 after Peek; got %d", count)
	}
}
