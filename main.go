package main

import (
	"fmt"

	"github.com/CB101b/go-queue/pkg/queue"
)

func main() {
	queue := new(queue.Queue[string])
	fmt.Println(queue)
	queue.Push("a")
	fmt.Println(queue)
}
