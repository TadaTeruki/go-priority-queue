package main

import (
	"fmt"

	priority_queue "github.com/TadaTeruki/go-priority-queue"
)

func main() {
	pq := new(priority_queue.PriorityQueue[string])

	pq.Insert("a", 7.5)
	pq.Insert("b", -2.1)
	pq.Insert("c", 3.2)
	pq.Insert("d", 0.0)

	item, priority, err := pq.Front()
	if err != nil {
		panic(err)
	}

	fmt.Println(item, priority)
}
