package stack

type Stack struct {
	items []string
}

// IsEmpty returns true if the stack is empty
func (s *Stack) IsEmpty() (bool, int) {
	nItems := len(s.items)

	if nItems == 0 {
		return true, 0
	}

	return false, nItems
}

// Push adds an item to the stack
func (s *Stack) Push(item string) {
	s.items = append(s.items, item)
}

// Pop removes the top item from the stack and returns it
func (s *Stack) Pop() (string, bool) {
	empty, nItems := s.IsEmpty()

	if empty {
		return "", false
	}

	top := s.items[nItems-1]
	s.items = s.items[:nItems-1]

	return top, true
}

// Peek returns the top item without removing it
func (s *Stack) Peek() (string, bool) {
	empty, nItems := s.IsEmpty()

	if empty {
		return "", false
	}

	top := s.items[nItems-1]

	return top, true
}
