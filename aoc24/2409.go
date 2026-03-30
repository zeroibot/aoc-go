package aoc24

import (
	. "github.com/zeroibot/aoc-go/aoc"
)

func Day09() Solution {
	numbers := data09(true)

	// Part 1
	head, tail := buildLinkedList(numbers)
	a, b := head, tail
	for {
		for a.value != nil {
			a = a.next
		}
		for b.value == nil {
			b = b.prev
		}
		if a.rank > b.rank {
			break
		}

		if a.size == b.size {
			a.value, b.value = b.value, a.value
		} else if a.size < b.size {
			a.value = b.value
			b.size = b.size - a.size
		} else if a.size > b.size {
			a.value, b.value = b.value, a.value
			free := a.size - b.size
			a.size = b.size
			node := newNode(nil, free, a.rank)
			insertNodeAfter(node, a)
		}
	}
	checksum1 := getChecksum(head)

	// Part 2
	head, tail = buildLinkedList(numbers)
	curr := tail
	for {
		fid := curr.value
		free := head
		for {
			if free != nil && free.value == nil && free.size >= curr.size {
				break
			}
			if free != nil && free.isTail {
				free = nil
				break
			} else {
				free = free.next
			}
		}

		if free != nil && free.rank < curr.rank {
			free.value, curr.value = curr.value, free.value
			if free.size != curr.size {
				left := free.size - curr.size
				free.size = curr.size
				node := newNode(nil, left, free.rank)
				insertNodeAfter(node, free)
			}
		}

		goal := 0
		if fid != nil {
			goal = *fid - 1
		}
		if goal > 0 {
			for {
				if curr.value != nil && *curr.value == goal {
					break
				}
				curr = curr.prev
			}
		} else {
			break
		}
	}
	checksum2 := getChecksum(head)

	return NewSolution(checksum1, checksum2)
}

func data09(full bool) []int {
	return ToIntLine(ReadFirstLine(24, 9, full))
}

type Node struct {
	value  *int
	size   int
	rank   int
	isHead bool
	isTail bool
	next   *Node
	prev   *Node
}

func newNode(value *int, size, rank int) *Node {
	node := &Node{value: value, size: size, rank: rank}
	node.next = node
	node.prev = node
	return node
}

func insertNodeAfter(node *Node, prev *Node) {
	node.prev = prev
	node.next = prev.next
	prev.next.prev = node
	prev.next = node
}

func buildLinkedList(numbers []int) (*Node, *Node) {
	count := len(numbers)
	zero := 0
	head := newNode(&zero, 0, 0)
	tail := newNode(&zero, 0, 0)
	var curr *Node = nil
	rank := 0
	for i := 0; i < count; i += 2 {
		fid := i / 2
		node1 := newNode(&fid, numbers[i], rank)
		rank += 1
		if curr == nil {
			head = node1
			head.isHead = true
		} else {
			curr.next = node1
			node1.prev = curr
		}

		if i < count-1 {
			if numbers[i+1] > 0 {
				node2 := newNode(nil, numbers[i+1], rank)
				rank += 1
				node1.next = node2
				node2.prev = node1
				curr = node2
			} else {
				curr = node1
			}
		} else {
			tail = node1
			tail.isTail = true
		}
	}
	return head, tail
}

func getChecksum(head *Node) int {
	checksum, i := 0, 0
	node := head
	for {
		for range node.size {
			if node.value != nil {
				checksum += (*node.value * i)
			}
			i += 1
		}
		if node.isTail {
			break
		}
		node = node.next
	}
	return checksum
}
