package main

import (
	sortPkg "ds/sort"
	"log"
)

func main() {
	s := sortPkg.Stack[int]{}
	s.Push(5)
	s.Push(7)
	s.Push(9)
	s.Push(1)
	s.Push(8)
	log.Println(s)
	log.Println(*s.Top.Prev.Prev.Prev)
}
