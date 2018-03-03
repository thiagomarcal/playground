package collections

// Queue Type
type Queue []interface{}

// Push Element
func (q *Queue) Push(n interface{}) {
	*q = append(*q, n)
}

// Pop Element
func (q *Queue) Pop() (n interface{}) {
	n = (*q)[0]
	*q = (*q)[1:]
	return
}

// Peek Queue
func (q *Queue) Peek() (n interface{}) {
	return (*q)[0]
}

// Len Queue
func (q *Queue) Len() int {
	return len(*q)
}
