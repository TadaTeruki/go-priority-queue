/*

MIT License

Copyright (c) 2022 Teruki TADA

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

*/

// Package priority_queue provides simple implementation of priority queue for golang.
package priority_queue

import "errors"

// PriorityQueue holds items within the structure of priority queue.
type PriorityQueue[T any] struct {
	Tree []item[T]
}

type item[T any] struct {
	item     T
	priority float64
}

// Size returns the number of items.
func (pq *PriorityQueue[T]) Size() int {
	return len(pq.Tree)
}

// Front returns the most prior item with its priority in PriorityQueue.
// If the PriorityQueue has no data, Front alternatively returns an error.
func (pq *PriorityQueue[T]) Front() (T, float64, error) {
	if pq.Size() == 0 {
		return *new(T), 0., errors.New("priority queue has no data")
	}

	obj := pq.Tree[0]

	return obj.item, obj.priority, nil
}

// Pop removes the most prior item from PriorityQueue.
func (pq *PriorityQueue[T]) Pop() {

	obj_id := 0
	pq.Tree[obj_id] = pq.Tree[pq.Size()-1]

	for obj_id < pq.Size() {
		a_id := obj_id*2 + 1
		b_id := a_id + 1

		a_ok := pq.Size() > a_id && pq.Tree[a_id].priority > pq.Tree[obj_id].priority
		b_ok := pq.Size() > b_id && pq.Tree[b_id].priority > pq.Tree[obj_id].priority

		if b_ok && (!a_ok || pq.Tree[b_id].priority > pq.Tree[a_id].priority) {
			swap := pq.Tree[obj_id]
			pq.Tree[obj_id] = pq.Tree[b_id]
			pq.Tree[b_id] = swap
			obj_id = b_id
		} else if a_ok {
			swap := pq.Tree[obj_id]
			pq.Tree[obj_id] = pq.Tree[a_id]
			pq.Tree[a_id] = swap
			obj_id = a_id
		} else {
			break
		}
	}
	pq.Tree = pq.Tree[0 : pq.Size()-1]
}

// Insert inserts a new item with its priority to PriorityQueue.
func (pq *PriorityQueue[T]) Insert(it T, priority float64) {
	pq.Tree = append(pq.Tree, item[T]{item: it, priority: priority})

	obj_id := pq.Size() - 1

	for obj_id != 0 {
		parent_id := (obj_id - 1) / 2
		if pq.Tree[obj_id].priority > pq.Tree[parent_id].priority {
			swap := pq.Tree[obj_id]
			pq.Tree[obj_id] = pq.Tree[parent_id]
			pq.Tree[parent_id] = swap
			obj_id = parent_id
		} else {
			break
		}
	}
}
