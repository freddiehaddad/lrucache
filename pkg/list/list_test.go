package list

import "testing"

func checkNodes(t *testing.T, list *List, nodes []*Node) {
	head := list.head
	for i := range nodes {
		if head != nodes[i] {
			t.Errorf("test[%d] failed. expected node=%+v got=%+v", i, nodes[i], head)
		}
		head = head.next
	}

	tail := list.tail
	for i := range nodes {
		if tail != nodes[len(nodes)-1-i] {
			t.Errorf("test[%d] failed. expected node=%+v got=%+v", i, nodes[i], tail)
		}
		tail = tail.prev
	}
}

func TestNew(t *testing.T) {
	list := New()
	nodes := []*Node{list.head, list.tail}
	checkNodes(t, list, nodes)
}

func TestEnqueue(t *testing.T) {
	list := New()
	nodes := []*Node{list.head, list.tail}
	checkNodes(t, list, nodes)

	for i := 0; i < 5; i++ {
		node := &Node{}
		list.Enqueue(node)
		nodes = append(nodes, nodes[len(nodes)-1])
		nodes[len(nodes)-2] = node
		checkNodes(t, list, nodes)
	}
}

func TestDequeue(t *testing.T) {
	list := New()
	nodes := []*Node{list.head, list.tail}
	checkNodes(t, list, nodes)

	for i := 0; i < 5; i++ {
		node := &Node{}
		list.Enqueue(node)
		nodes = append(nodes, nodes[len(nodes)-1])
		nodes[len(nodes)-2] = node
		checkNodes(t, list, nodes)
	}

	for len(nodes)-2 > 0 {
		dequeued := list.Dequeue()
		expected := nodes[1]
		if dequeued != expected {
			t.Errorf("Dequeue failed. expected=%+v got=%+v", expected, dequeued)
		}
		copy(nodes[1:], nodes[2:])
		nodes = nodes[:len(nodes)-1]
		checkNodes(t, list, nodes)
	}
}

func TestDelete(t *testing.T) {
	tests := []struct {
		nodes   []*Node
		deletes []int
	}{
		{nodes: make([]*Node, 3), deletes: []int{1}},
		{nodes: make([]*Node, 4), deletes: []int{1, 1}},
		{nodes: make([]*Node, 4), deletes: []int{2, 1}},
		{nodes: make([]*Node, 5), deletes: []int{2, 1, 1}},
		{nodes: make([]*Node, 5), deletes: []int{2, 2, 1}},
		{nodes: make([]*Node, 6), deletes: []int{1, 3, 2, 1}},
		{nodes: make([]*Node, 6), deletes: []int{2, 2, 1, 1}},
	}

	for _, test := range tests {
		list := New()
		test.nodes[0] = list.head
		idx := 1
		for j := 0; j < len(test.deletes); j++ {
			node := &Node{Value: idx}
			test.nodes[idx] = node
			list.Enqueue(node)
			idx++
		}
		test.nodes[idx] = list.tail
		checkNodes(t, list, test.nodes)

		for j := 0; j < len(test.deletes); j++ {
			idx = test.deletes[j]
			list.Delete(test.nodes[idx])
			copy(test.nodes[idx:], test.nodes[idx+1:])
			test.nodes = test.nodes[:len(test.nodes)-1]
			checkNodes(t, list, test.nodes)
		}
	}
}
