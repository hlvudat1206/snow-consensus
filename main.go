// Purpose
// 1. Initializes a network of 200 nodes.
// 2. Simulates consensus by running all nodes in parallel.
// 3. Waits for all nodes to complete and checks for consensus.

package main

import (
	"fmt"
	"snow-consensus/snow-consensus/node"

	// "snow-consensus/node"
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
		// Starts the consensus process for each node in parallel using goroutines.
		go func(n *node.Node) {
			defer wg.Done()
			n.Start()
		}(n)
	}

	wg.Wait()
	fmt.Println("Consensus reached!")
}
