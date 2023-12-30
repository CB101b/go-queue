package main

import (
	"fmt"

	"github.com/CB101b/go-queue/pkg/queue"
)

func main() {
	queue := new(queue.Queue[string])
	fmt.Println(queue)
	queue.Push("a")
	queue.Push("b")
	fmt.Println(queue)
	fmt.Println(queue.Take())
	fmt.Println(queue)
}
