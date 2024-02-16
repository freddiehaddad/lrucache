// Package list provides a minimal doubly linked list that exposes its
// internal data structures for packages needing direct access.
//
// It's not meant to be a general purpose lnked list.
package list

// Node is an internal
type Node struct {
	next  *Node
	prev  *Node
	Key   int
	Value int
}

type List struct {
	head *Node
	tail *Node
}

// New returns an empty linked list.
func New() *List {
	l := &List{
		head: &Node{},
		tail: &Node{},
	}

	l.head.next = l.tail
	l.tail.prev = l.head

	return l
}

// Enqueue adds node to the end of the list.
func (l *List) Enqueue(n *Node) {
	prev := l.tail.prev
	prev.next = n
	n.prev = prev
	l.tail.prev = n
	n.next = l.tail
}

// Dequeue removes the first elememnt in the list and returns it.
// No bounds checking happens.  Calling this function on an empty list
// will result in undefined behavior.
func (l *List) Dequeue() *Node {
	node := l.head.next
	l.Delete(node)
	return node
}

// Delete removes node from the list by linkin it's previous and
// next pointers. Calling this function with an invalid or nil object
// will result in undefined behavior.
func (l *List) Delete(node *Node) {
	prev := node.prev
	next := node.next
	prev.next = next
	next.prev = prev
}
