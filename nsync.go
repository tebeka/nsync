// Package nsync provides type safe wrappers around sync.Pool and sync.Map.
package nsync

import "sync"

// Pool is a type safe wrapper around sync.Pool
// See sync.Pool documentation for full explanation.
type Pool[T any] struct {
	pool sync.Pool
}

// NewPool return a new pool. newFn optionally specifies a function to generate
// a value when Get would otherwise return nil.
func NewPool[T any](newFn func() T) *Pool[T] {
	var p Pool[T]
	if newFn != nil {
		p.pool.New = func() any {
			return newFn()
		}
	}

	return &p
}

// Put puts v in the pool.
func (p *Pool[T]) Put(v T) {
	p.pool.Put(v)
}

// Get gets an item from the pool.
// If there are not items in the pool, the second return value will be false.
func (p *Pool[T]) Get() (T, bool) {
	v := p.pool.Get()
	if v == nil {
		var zero T
		return zero, false
	}
	return v.(T), true
}

// Map is a type safe wrapper around sync.Map.
// See sync.Map documentation for full explanation.
type Map[K comparable, V any] struct {
	m sync.Map
}

// CompareAndDelete deletes the entry for key if its value is equal to old.
func (m *Map[K, V]) CompareAndDelete(key K, old V) (deleted bool) {
	return m.m.CompareAndDelete(key, old)
}

// CompareAndSwap swaps the old and new values for key if the value stored in the map is equal to old.
func (m *Map[K, V]) CompareAndSwap(key K, old, new V) bool {
	return m.m.CompareAndSwap(key, old, new)
}

// Delete deletes the value for a key.
func (m *Map[K, V]) Delete(key K) {
	m.m.Delete(key)
}

// Load returns the value stored in the map for a key, or nil if no value is present.
func (m *Map[K, V]) Load(key K) (value V, ok bool) {
	v, ok := m.m.Load(key)
	if !ok {
		var zero V
		return zero, ok
	}

	return v.(V), ok
}

// LoadAndDelete deletes the value for a key, returning the previous value if any.
func (m *Map[K, V]) LoadAndDelete(key K) (value V, loaded bool) {
	v, loaded := m.m.Load(key)
	if !loaded {
		var zero V
		return zero, loaded
	}

	return v.(V), loaded
}

// LoadOrStore returns the existing value for the key if present.
// Otherwise, it stores and returns the given value.
// The loaded result is true if the value was loaded, false if stored.
func (m *Map[K, V]) LoadOrStore(key K, value V) (actual V, loaded bool) {
	v, loaded := m.m.LoadOrStore(key, value)
	return v.(V), loaded

}

// Range calls f sequentially for each key and value present in the map.
// If f returns false, range stops the iteration.
func (m *Map[K, V]) Range(f func(key K, value V) bool) {
	fn := func(k, v any) bool {
		return f(k.(K), v.(V))
	}
	m.m.Range(fn)
}

// Store sets the value for a key.
func (m *Map[K, V]) Store(key K, value V) {
	m.m.Store(key, value)
}

// Swap swaps the value for a key and returns the previous value if any.
// The loaded result reports whether the key was present.
func (m *Map[K, V]) Swap(key K, value V) (previous V, loaded bool) {
	p, loaded := m.m.Swap(key, value)
	if !loaded {
		var zero V
		return zero, loaded
	}
	return p.(V), loaded
}
