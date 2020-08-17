package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func main() {
	fmt.Println("Go Tour - equivalent binary tree")
	t1 := tree.New(1)
	t2 := tree.New(1)

	// This is how to print tree
	fmt.Print("Tree 1: ")
	ch := make(chan int)
	go func() {
		walk(t1, ch)
		close(ch)
	}()
	for v := range ch {
		fmt.Printf("%d ", v)
	}
	fmt.Println()

	fmt.Print("Tree 2: ")
	ch = make(chan int)
	go func() {
		walk(t2, ch)
		close(ch)
	}()
	for v := range ch {
		fmt.Printf("%d ", v)
	}
	fmt.Println()

	// Check whether tree is equivalent
	fmt.Println(same(t1, t2))
}

func walk(t *tree.Tree, ch chan int) {
	if t != nil {
		walk(t.Left, ch)
		ch <- t.Value
		walk(t.Right, ch)
	}
}

func same(t1 *tree.Tree, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	res := make(chan bool, 1)
	go walk(t1, ch1)
	go walk(t2, ch2)
	go func() {
		for i := 0; i < 10; i++ {
			v1, v2 := <-ch1, <-ch2
			if v1 != v2 {
				res <- false
				return
			}
		}
		res <- true
	}()
	return <-res
}
