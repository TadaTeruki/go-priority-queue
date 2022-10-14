# go-priority-queue
Simple implementation of Priority Queue in Golang.

## Usage

```go

package main

import (
	"fmt"

	priority_queue "github.com/TadaTeruki/go-priority-queue"
)

func main() {
	pq := new(priority_queue.PriorityQueue[string])

	pq.Push("a", 7.5)
	pq.Push("b", -2.1)
	pq.Push("c", 3.2)
	pq.Push("d", 0.0)

	item, priority, err := pq.Front()
	if err != nil {
		panic(err)
	}

	fmt.Println(item, priority)
}
```

## Installation

```
 $ go get -u github.com/TadaTeruki/go-priority-queue
```

## LICENSE

MIT License<br>
Copyright (c) 2022 Teruki TADA
