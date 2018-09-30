package script

// A Stack is a FIFO data structure which orders data.
type stack [][]byte

// Push appends a given value to the top of the stack.
//
// The nil value is ignored.
func (s *stack) Push(v []byte) {
	if v == nil {
		return
	}

	*s = append(*s, v)
}

// Pop yields the top of the stack.
//
// If the stack is empty nil is returned.
func (s *stack) Pop() []byte {
	if len(*s) == 0 {
		return nil
	}

	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]

	return v
}

// Size returns the number of items on the stack.
func (s *stack) Size() int {
	return len(*s)
}
