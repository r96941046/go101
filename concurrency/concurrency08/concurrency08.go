package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

/*
   type Tree struct {
       Left  *Tree
       Value int
       Right *Tree
   }
*/

// Walk walks the tree t sending all values
// from the tree to the channel ch
func Walk(t *tree.Tree, ch chan int) {
	_walk(t, ch)
	close(ch)
}

func _walk(t *tree.Tree, ch chan int) {
	if t != nil {
		_walk(t.Left, ch)
		ch <- t.Value
		_walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values
func Same(t1, t2 *tree.Tree) bool {
	c1 := make(chan int)
	c2 := make(chan int)

	go Walk(t1, c1)
	go Walk(t2, c2)

	for i := range c1 {
		if i != <-c2 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
