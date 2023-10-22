package nsync_test

import (
	"fmt"
	"sort"

	"github.com/tebeka/nsync"
)

func ExamplePool() {
	var counter int
	fn := func() int {
		counter++
		return counter
	}
	pool := nsync.NewPool[int](fn)
	fmt.Println(pool.Get())
	pool.Put(3)
	fmt.Println(pool.Get())

	// Pool without New function.
	pool = nsync.NewPool[int](nil)
	fmt.Println(pool.Get())

	// Output:
	// 1 true
	// 3 true
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
