package collections

// Stack type
type Stack []interface{}

// Push Stack
func (q *Stack) Push(n interface{}) {
	*q = append(*q, n)
}

// Pop Stack
func (q *Stack) Pop() (n interface{}) {
	x := q.Len() - 1
	n = (*q)[x]
	*q = (*q)[:x]
	return
}

// Peek Stack
func (q *Stack) Peek() interface{} {
	return (*q)[q.Len()-1]
}

// Len Stack
func (q *Stack) Len() int {
	return len(*q)
}
