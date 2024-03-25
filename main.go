package main

import (
	"ds/array"
	sortPkg "ds/sort"
	"log"
)

func main() {
	log.Println(array.LinearSearch([]int{1, 2, 3, 4, 5}, 3))
	log.Println(array.BinarySearch([]int{1, 2, 3, 4, 5}, 3))
	sortPkg.BubbleSort([]int{3, 4, 45, 6, 2, 1, 234, 5, 2, 1, 2})

	q := sortPkg.Queue[int]{}

	q.Enqueue(1)
	q.Enqueue(4)
	q.Enqueue(6)
	q.Enqueue(3)
	log.Println(q.Peek())
	log.Println(q.Tail)
	log.Println(q.Tail.Value)
}
