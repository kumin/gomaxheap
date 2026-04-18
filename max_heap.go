package maxheap

import (
	"log/slog"
	"slices"
)

type Node struct {
	Seed   any
	Weight float64
}

type Heap struct {
	nodes []*Node
}

func NewHeap(nodes []*Node) *Heap {
	heap := &Heap{nodes: nodes}
	buildHeap(heap.nodes)

	return heap
}

func (h *Heap) Pop() *Node {
	if h.Size() == 0 {
		slog.Warn("heap is empty")
		return nil
	}
	root := h.nodes[0]

	h.nodes[0] = h.nodes[h.Size()-1]
	h.nodes = h.nodes[:h.Size()-1]

	parentIdx := 0
	nextParentIdx := parentIdx
	n := h.Size()
	for parentIdx*2+1 < n {
		if parentIdx*2+1 < n && h.nodes[parentIdx].Weight < h.nodes[parentIdx*2+1].Weight {
			h.nodes[parentIdx], h.nodes[parentIdx*2+1] = h.nodes[parentIdx*2+1], h.nodes[parentIdx]
			nextParentIdx = parentIdx*2 + 1
		}
		if parentIdx*2+2 < n && h.nodes[parentIdx].Weight < h.nodes[parentIdx*2+2].Weight {
			h.nodes[parentIdx], h.nodes[parentIdx*2+2] = h.nodes[parentIdx*2+2], h.nodes[parentIdx]
			nextParentIdx = parentIdx*2 + 2
		}
		parentIdx = nextParentIdx
	}

	return root
}

func (h *Heap) Insert(node *Node) {
	h.nodes = append(h.nodes, node)
	n := h.Size()
	i := n - 1
	for i > 0 {
		parentIdx := getParentIdx(i)
		if parentIdx*2+1 < n && h.nodes[parentIdx].Weight < h.nodes[parentIdx*2+1].Weight {
			h.nodes[parentIdx], h.nodes[parentIdx*2+1] = h.nodes[parentIdx*2+1], h.nodes[parentIdx]
		}
		if parentIdx*2+2 < n && h.nodes[parentIdx].Weight < h.nodes[parentIdx*2+2].Weight {
			h.nodes[parentIdx], h.nodes[parentIdx*2+2] = h.nodes[parentIdx*2+2], h.nodes[parentIdx]
		}

		i = parentIdx
	}
}

func (h *Heap) Delete(nodeIdx int) {

}

func (h *Heap) Size() int {
	return len(h.nodes)
}

func (h *Heap) GetNodes() []*Node {
	return slices.Clone(h.nodes)
}

func buildHeap(nodes []*Node) {
	n := len(nodes)
	i := n - 1
	for i > 1 {
		parentIdx := getParentIdx(i)
		heapify(parentIdx, nodes)
		i -= 2
	}
}

func (h *Heap) PrintHeap() {
	slog.Info("=========HEAP NODES==========")
	for i, node := range h.nodes {
		slog.Info("node", slog.Int("idx", i), slog.Any("value", node))
	}
}

func heapify(parentIdx int, nodes []*Node) {
	n := len(nodes)
	if parentIdx*2+1 < n && nodes[parentIdx].Weight < nodes[parentIdx*2+1].Weight {
		nodes[parentIdx], nodes[parentIdx*2+1] = nodes[parentIdx*2+1], nodes[parentIdx]
		heapify(parentIdx*2+1, nodes)
	}
	if parentIdx*2+2 < n && nodes[parentIdx].Weight < nodes[parentIdx*2+2].Weight {
		nodes[parentIdx], nodes[parentIdx*2+2] = nodes[parentIdx*2+2], nodes[parentIdx]
		heapify(parentIdx*2+2, nodes)
	}
}

func getParentIdx(idx int) int {
	if idx < 1 {
		return idx
	}
	var parentIdx int
	if idx%2 == 0 {
		parentIdx = (idx - 2) / 2
	} else {
		parentIdx = (idx - 1) / 2
	}

	return parentIdx
}
