package main

import (
	"fmt"

	pq "github.com/TadaTeruki/go-priority-queue/v2"
)

func main() {
	priorityQueue := new(pq.PriorityQueue[string])

	priorityQueue.Push("a", 7.5)
	priorityQueue.Push("b", -2.1)
	priorityQueue.Push("c", 3.2)
	priorityQueue.Push("d", 0.0)

	item, priority := priorityQueue.Pop()

	fmt.Println(item, priority)
}
