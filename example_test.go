package nsync_test

import (
	"fmt"
	"sort"

	"github.com/tebeka/nsync"
)

var counter int

func ExamplePool() {
	pool := nsync.Pool[int]{
		New: func() int {
			counter++
			return counter
		},
	}
	fmt.Println(pool.Get())
	pool.Put(2)
	fmt.Println(pool.Get())

	pool = nsync.Pool[int]{}
	fmt.Println(pool.Get())

	// Output:
	// 1 true
	// 2 true
	// 0 false
}

func ExampleMap() {
	var m nsync.Map[string, int]
	m.Store("who", 1)
	m.Store("what", 2)
	fmt.Println(m.Load("who"))

	var items []string
	m.Range(func(key string, value int) bool {
		items = append(items, fmt.Sprintf("%s -> %d", key, value))
		return true
	})
	// Sort to get consistent prints.
	sort.Strings(items)
	for _, i := range items {
		fmt.Println(i)
	}

	// Output:
	// 1 true
	// what -> 2
	// who -> 1
}
