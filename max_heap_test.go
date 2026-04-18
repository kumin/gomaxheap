package maxheap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeapify(t *testing.T) {
	testCases := []struct {
		name  string
		nodes []*Node
		heap  []*Node
	}{
		{
			name:  "test case 1",
			nodes: []*Node{{Seed: "kumin", Weight: 2}, {Seed: "midu", Weight: 10}, {Seed: "omachi", Weight: 3}, {Seed: "may", Weight: 6}, {Seed: "lulu", Weight: 4}},
			heap:  []*Node{{Seed: "midu", Weight: 10}, {Seed: "may", Weight: 6}, {Seed: "omachi", Weight: 3}, {Seed: "kumin", Weight: 2}, {Seed: "lulu", Weight: 4}},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			heap := NewHeap(tt.nodes)
			heap.PrintHeap()
			assert.Equal(t, heap.GetNodes(), tt.heap)
		})
	}
}

func TestInsert(t *testing.T) {
	type args struct {
		nodes        []*Node
		insertedNode *Node
		heap         []*Node
	}

	testCases := []struct {
		name string
		args *args
	}{
		{
			name: "test case 1",
			args: &args{
				nodes: []*Node{{Seed: "kumin", Weight: 2}, {Seed: "midu", Weight: 10}, {Seed: "omachi", Weight: 3}, {Seed: "may", Weight: 6}, {Seed: "lulu", Weight: 4}},
				insertedNode: &Node{
					Seed:   "nhan",
					Weight: 100,
				},
				heap: []*Node{{Seed: "nhan", Weight: 100}, {Seed: "may", Weight: 6}, {Seed: "midu", Weight: 10}, {Seed: "kumin", Weight: 2}, {Seed: "lulu", Weight: 4}, {Seed: "omachi", Weight: 3}},
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			heap := NewHeap(tt.args.nodes)
			heap.Insert(tt.args.insertedNode)
			heap.PrintHeap()
			assert.Equal(t, heap.GetNodes(), tt.args.heap)
		})
	}
}

func TestPop(t *testing.T) {
	type args struct {
		nodes        []*Node
		insertedNode *Node
	}

	testCases := []struct {
		name string
		args *args
	}{
		{
			name: "test case 1",
			args: &args{
				nodes: []*Node{{Seed: "kumin", Weight: 2}, {Seed: "midu", Weight: 10}, {Seed: "omachi", Weight: 3}, {Seed: "may", Weight: 6}, {Seed: "lulu", Weight: 4}},
				insertedNode: &Node{
					Seed:   "nhan",
					Weight: 100,
				},
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			heap := NewHeap(tt.args.nodes)
			heap.Insert(tt.args.insertedNode)
			assert.Equal(t, heap.Pop(), tt.args.insertedNode)
			heap.PrintHeap()
		})
	}
}
