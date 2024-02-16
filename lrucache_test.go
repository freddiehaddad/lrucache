package lrucache

import (
	"testing"
)

type action struct {
	action   string
	key      int
	value    int
	expected int
}

func TestNew(t *testing.T) {
	tests := []struct {
		capacity int
	}{
		{1},
		{2},
		{3},
	}

	for i, test := range tests {
		cache := New(test.capacity)

		if len(cache.memory) != test.capacity {
			t.Errorf("test[%d] expected capcity=%d got=%d", i, test.capacity, len(cache.memory))
		}

		// dummy nodes
		for c := 1; c <= test.capacity; c++ {
			node := cache.list.Dequeue()
			if node.Key != -c {
				t.Errorf("test[%d] key wrong for node expected=%d got=%d", i, -c, node.Key)
			}

			if node.Value != -c {
				t.Errorf("test[%d] value wrong for node expected=%d got=%d", i, -c, node.Value)
			}

			if _, ok := cache.memory[-c]; !ok {
				t.Errorf("test[%d] key=%d not found", i, -c)
			}
		}
	}
}

func TestPut(t *testing.T) {
	tests := []struct {
		capacity int
		puts     []action   // put actions
		expected [][]action // cache layout after put
	}{
		{
			1,
			[]action{
				{key: 1, value: 1},
				{key: 1, value: 2},
				{key: 2, value: 1},
				{key: 2, value: 2},
			},
			[][]action{
				{{key: 1, value: 1}},
				{{key: 1, value: 2}},
				{{key: 2, value: 1}},
				{{key: 2, value: 2}},
			},
		},
		{
			2,
			[]action{
				{key: 1, value: 1},
				{key: 2, value: 1},
				{key: 1, value: 2},
				{key: 3, value: 1},
			},
			[][]action{
				{{key: -2, value: -2}, {key: 1, value: 1}},
				{{key: 1, value: 1}, {key: 2, value: 1}},
				{{key: 2, value: 1}, {key: 1, value: 2}},
				{{key: 1, value: 2}, {key: 3, value: 1}},
			},
		},
		{
			3,
			[]action{
				{key: 1, value: 1},
				{key: 2, value: 1},
				{key: 1, value: 2},
				{key: 3, value: 1},
				{key: 2, value: 2},
				{key: 4, value: 1},
			},
			[][]action{
				{{key: -2, value: -2}, {key: -3, value: -3}, {key: 1, value: 1}},
				{{key: -3, value: -3}, {key: 1, value: 1}, {key: 2, value: 1}},
				{{key: -3, value: -3}, {key: 2, value: 1}, {key: 1, value: 2}},
				{{key: 2, value: 1}, {key: 1, value: 2}, {key: 3, value: 1}},
				{{key: 1, value: 2}, {key: 3, value: 1}, {key: 2, value: 2}},
				{{key: 3, value: 1}, {key: 2, value: 2}, {key: 4, value: 1}},
			},
		},
	}

	for tidx, test := range tests {
		cache := New(test.capacity)

		for pidx := range test.puts {
			put := test.puts[pidx]
			cache.Put(put.key, put.value)

			if len(cache.memory) != test.capacity {
				t.Errorf("test[%d] expected capcity=%d got=%d", tidx, test.capacity, len(cache.memory))
			}

			for eidx, expected := range test.expected[pidx] {
				if node, ok := cache.memory[expected.key]; !ok {
					t.Errorf("test[%d] key=%d value=%d not in cache", eidx, expected.key, expected.value)
				} else {
					if node.Key != expected.key {
						t.Errorf("test[%d] node key wrong expected=%d got=%d", tidx, expected.key, node.Key)
					}
					if node.Value != expected.value {
						t.Errorf("test[%d] node value wrong expected=%d got=%d", tidx, expected.value, node.Value)
					}
				}
			}

		}
	}
}

func TestGet(t *testing.T) {
	tests := []struct {
		capacity int
		actions  []action
	}{
		{
			1,
			[]action{
				{"get", 1, 0, -1},
				{"put", 1, 1, 0},
				{"get", 1, 0, 1},
				{"put", 1, 2, 0},
				{"get", 1, 0, 2},
				{"put", 2, 1, 0},
				{"get", 1, 0, -1},
			},
		},
		{
			2,
			[]action{
				{"put", 1, 1, 0},
				{"get", 1, 0, 1},
				{"put", 2, 1, 0},
				{"get", 1, 0, 1},
				{"get", 2, 0, 1},
				{"put", 1, 2, 0},
				{"get", 1, 0, 2},
				{"put", 3, 1, 0},
				{"get", 2, 0, -1},
				{"get", 3, 0, 1},
				{"put", 2, 1, 0},
				{"get", 2, 0, 1},
			},
		},
		{
			3,
			[]action{
				{"put", 1, 1, 0},
				{"put", 2, 1, 0},
				{"put", 3, 1, 0},
				{"get", 3, 0, 1},
				{"get", 2, 0, 1},
				{"get", 1, 0, 1},
				{"put", 4, 1, 0},
				{"get", 3, 0, -1},
				{"get", 2, 0, 1},
				{"get", 1, 0, 1},
			},
		},
	}

	for tidx, test := range tests {
		cache := New(test.capacity)
		for aidx, action := range test.actions {
			switch action.action {
			case "put":
				cache.Put(action.key, action.value)
			case "get":
				if result := cache.Get(action.key); result != action.expected {
					t.Errorf("test[%d] action[%d] get failed. expected=%d got=%d", tidx, aidx, action.expected, result)
				}
			default:
				t.Errorf("test[%d] action[%d] unknown action=%q", tidx, aidx, action.action)
			}
		}
	}
}
