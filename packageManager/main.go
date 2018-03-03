package main

import (
	"fmt"

	"github.com/oleiade/lane"
)

func main() {

	queue := lane.NewQueue()
	queue.Enqueue("grumpyClient")
	queue.Enqueue("happyClient")
	queue.Enqueue("ecstaticClient")

	fmt.Println(queue.Pop())

}
