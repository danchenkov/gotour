package main

import (
	"fmt"

	"github.com/fatih/color"
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	Walk(t.Right, ch)
	ch <- t.Value
}

func WalkWrapper(t *tree.Tree, ch chan int) {
	Walk(t, ch)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	match := make(map[int]int)
	ch1 := make(chan int)
	ch2 := make(chan int)
	go WalkWrapper(t1, ch1)
	go WalkWrapper(t2, ch2)

	for {
		v1, ok1 := <-ch1
		if !ok1 {
			break
		}
		match[v1] += 1
	}
	for {
		v2, ok2 := <-ch2
		if !ok2 {
			break
		}
		match[v2] += 1
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
	go WalkWrapper(tree.New(1), ch)

	for i := 0; i < 10; i++ {
		v := <-ch
		fmt.Printf("%d ", v)
	}

	fmt.Println()

	fmt.Print("tree.New(1) == tree.New(1): ")
	if Same(tree.New(1), tree.New(1)) {
		color.Green("PASSED")
	} else {
		color.Red("FAILED")
	}

	fmt.Print("tree.New(1) != tree.New(2): ")
	if !Same(tree.New(1), tree.New(2)) {
		color.Green("PASSED")
	} else {
		color.Red("FAILED")
	}
}
