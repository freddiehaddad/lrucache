package lrucache

import (
	"github.com/freddiehaddad/lrucache/pkg/list"
)

// The least recently used (LRU) cache
type LRUCache struct {
	capacity int
	memory   map[int]*list.Node // the cache store
	list     *list.List
}

// New returns a new fully associative LRU (least recently used) with set
// to the fixed capacity.
func New(capacity int) LRUCache {
	l := LRUCache{
		capacity: capacity,
		memory:   make(map[int]*list.Node),
		list:     list.New(),
	}

	for i := 1; i <= capacity; i++ {
		node := newNode()
		l.store(node, -i, -i)
	}

	return l
}

// Get retrieves the value for key from the cache and returns it.  The
// accessed value is updated such that it's the most recently used.  -1 is
// returned when key does not exist in the cache.
func (l *LRUCache) Get(key int) int {
	// hit
	if node := l.exists(key); node != nil {
		l.update(node, node.Value)
		return node.Value
	}

	// miss
	return -1
}

// Put stores the value in the cache for reference by key.  If the key
// already is already in the cache, its updated such that it's the most
// recently used. Otherwise the key/value pair is added and set as the most
// recently used. Put operations will replace the least recently used entry
// when at capacity.
func (l *LRUCache) Put(key int, value int) {
	// hit
	if data := l.exists(key); data != nil {
		l.update(data, value)
		return
	}

	// miss
	data := l.evict()
	l.store(data, key, value)
}

// remove element d from the doubly linked list and sets the objects
// pointed to by d.prev and d.next such that they pointing to each other.
func (l *LRUCache) remove(node *list.Node) {
	l.list.Delete(node)
}

// insert adds d to the tail of the doubly linked list.
func (l *LRUCache) insert(node *list.Node) {
	l.list.Enqueue(node)
}

// evict removes the least recently used entry from the cache and returns
// the evicted element.
func (l *LRUCache) evict() *list.Node {
	node := l.list.Dequeue()
	delete(l.memory, node.Key)
	return node
}

// store the key/value pair in the cache
func (l *LRUCache) store(node *list.Node, key, value int) {
	node.Key = key
	node.Value = value
	l.memory[key] = node
	l.list.Enqueue(node)
}

// exists returns the data object for key if it exists in the cache.
// Otherwise nil is returned.
func (l *LRUCache) exists(key int) *list.Node {
	return l.memory[key]
}

// update data with new value and set as most recently used entry.
func (l *LRUCache) update(node *list.Node, value int) {
	node.Value = value
	l.remove(node)
	l.insert(node)
}

// newNode creates a zero-value data object and returns it.  It is used
// exclusively for initializing a new LRUCache.
func newNode() *list.Node {
	return &list.Node{}
}
