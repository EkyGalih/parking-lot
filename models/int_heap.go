package models

import (
	"container/heap"
	"fmt"
	"strings"
)

// IntHeap adalah min-heap integer, digunakan untuk melacak slot parkir kosong
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push menambahkan elemen baru ke heap
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }

// Pop menghapus dan mengembalikan elemen terkecil dari heap
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// NewIntHeap membuat heap dengan slot dari 1 sampai seterusnya
func NewIntHeap(size int) *IntHeap {
	h := &IntHeap{}
	for i := 1; i <= size; i++ {
		*h = append(*h, i)
	}
	heap.Init(h)
	return h
}

func colorize(color, text string) string {
	colors := map[string]string{
		"black":   "\033[30m",
		"red":     "\033[31m",
		"green":   "\033[32m",
		"yellow":  "\033[33m",
		"blue":    "\033[34m",
		"magenta": "\033[35m",
		"cyan":    "\033[36m",
		"white":   "\033[37m",
		"reset":   "\033[0m",
	}

	code, ok := colors[strings.ToLower(color)]
	if !ok {
		return text
	}
	return fmt.Sprintf("%s%s%s", code, text, colors["reset"])
}
