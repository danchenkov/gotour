package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

const Total int = 10

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		go Walk(t.Left, ch)
	}
	if t.Right != nil {
		go Walk(t.Right, ch)
	}
	ch <- t.Value
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	match := make(map[int]int)
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for i := 0; i < Total*2; i++ {
		select {
		case v := <-ch1:
			match[v] += 1
		case v := <-ch2:
			match[v] += 1
		}
	}
	for _, v := range match {
		if v != 2 {
			return false
		}
	}
	return true
}

func main() {
	ch := make(chan int, 10)
	go Walk(tree.New(1), ch)

	fmt.Printf("tree.New(1) vs tree.New(1): Same? %t\n", Same(tree.New(1), tree.New(1)))
	fmt.Printf("tree.New(1) vs tree.New(2): Same? %t\n", Same(tree.New(1), tree.New(2)))
}
