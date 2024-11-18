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

	// create and start nodes
	for i := 0; i < TotalNodes; i++ {
		nodes[i] = node.NewNode(i, TotalNodes)
	}

	// create a channel to distribute nodes among workers
	nodeJobs := make(chan *node.Node, TotalNodes)

	// start 10 threads
	for t := 0; t < Processes; t++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for n := range nodeJobs {
				fmt.Printf("worker %d processing node %d\n", workerID, n.Id)
				n.Start() // simulate consensus process for this node
			}
		}(t)
	}

	for _, n := range nodes {
		nodeJobs <- n
	}
	close(nodeJobs)

	// wait for all workers to finish
	wg.Wait()

	// check if consensus was reached
	preference := nodes[0].Consensus.GetPreference()
	allAgreed := true
	for _, n := range nodes {
		if n.Consensus.GetPreference() != preference {
			allAgreed = false
			break
		}
	}

	if allAgreed {
		fmt.Printf("Consensus reached! All nodes agree on: %s\n", preference)
	} else {
		fmt.Println("Consensus was not reached by all nodes.")
	}
}
