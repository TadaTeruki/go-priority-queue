package pq_test

import (
	"testing"

	pq "github.com/TadaTeruki/go-priority-queue/v2"
)

func TestPriorityQueue(t *testing.T) {
	pq := new(pq.PriorityQueue[int]) // create a new priority queue

	if pq.Size() != 0 {
		t.Errorf("expected size of 0 but got %d", pq.Size())
	}

	// push some items into the priority queue
	pq.Push(3, 1.0)
	pq.Push(1, 2.0)
	pq.Push(4, 0.5)

	if pq.Size() != 3 {
		t.Errorf("expected size of 3 but got %d", pq.Size())
	}

	// check the front of the priority queue
	front, priority := pq.Front()
	if front == nil || *front != 1 || priority != 2.0 {
		t.Errorf("expected front item to be 1 with priority 2.0 but got %v with priority %f", front, priority)
	}

	// pop the most prior item from the priority queue
	item, priority := pq.Pop()
	if item == nil || *item != 1 || priority != 2.0 {
		t.Errorf("expected popped item to be 1 with priority 2.0 but got %v with priority %f", item, priority)
	}

	// check the size of the priority queue after pop
	if pq.Size() != 2 {
		t.Errorf("expected size of 2 but got %d", pq.Size())
	}
}
