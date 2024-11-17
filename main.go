package main

import (
	"fmt"
	"snow-consensus/node"
	"sync"
)

const (
	TotalNodes = 200
	Processes  = 10
)

func main() {
	var wg sync.WaitGroup
	nodes := make([]*node.Node, TotalNodes)

	// Create and start nodes
	for i := 0; i < TotalNodes; i++ {
		nodes[i] = node.NewNode(i, TotalNodes)
	}

	// Simulate node communication
	for _, n := range nodes {
		wg.Add(1)
		go func(n *node.Node) {
			defer wg.Done()
			n.Start()
		}(n)
	}

	wg.Wait()
	fmt.Println("Consensus reached!")
}
