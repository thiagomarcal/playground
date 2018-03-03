package main

import (
	"fmt"

	"github.com/thiagomarcal/playground/stackAndQueue/collections"
)

func main() {

	var queue collections.Queue
	var stack collections.Stack

	queue.Push([]string{"teste1", "teste2"})
	queue.Push(1)
	queue.Push("Thiago")

	stack.Push("Thiago")
	stack.Push([]string{"teste3", "teste4"})
	stack.Push([]int{1, 2})

	fmt.Println(queue)
	fmt.Println(stack)

	fmt.Printf("Queue Pop %v\n", queue.Pop())
	fmt.Printf("Stack Pop %v\n", stack.Pop())

}
