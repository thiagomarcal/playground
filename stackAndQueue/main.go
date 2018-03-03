package main

import (
	"fmt"

	"github.com/thiagomarcal/playground/stackAndQueue/collections"
)

func main() {

	var queue collections.Queue
	var stack collections.Stack

	queue.Push("Thiago")
	queue.Push("Daniel")
	queue.Push("Samoa")

	stack.Push("Thiago")
	stack.Push("Daniel")
	stack.Push("Samoa")

	fmt.Println(queue)
	fmt.Println(stack)

	fmt.Printf("Queue Pop %s\n", queue.Pop())
	fmt.Printf("Stack Pop %s\n", stack.Pop())

}
