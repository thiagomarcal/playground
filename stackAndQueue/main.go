package main

import (
	"fmt"

	"github.com/thiagomarcal/playground/stackAndQueue/collections"
)

func main() {

	var queue collections.Queue
	var stack collections.Stack

	queue.Push("Thiago")
	queue.Push("Mauro")
	queue.Push("Daniel")

	stack.Push("Thiago")
	stack.Push([]string{"teste3", "teste4"})
	stack.Push([]int{1, 2})

	fmt.Println(queue)
	fmt.Println(stack)

	fmt.Printf("Queue Pop %v\n", queue.Pop())
	fmt.Printf("Stack Pop %v\n", stack.Pop())

	fmt.Printf("Queue Peek %v\n", queue.Peek())
	fmt.Printf("Stack Peek %v\n", stack.Peek())

}
