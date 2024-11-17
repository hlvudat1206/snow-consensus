package tests

import (
	// "snow-consensus/node"
	"log"
	"snow-consensus/snow-consensus/node"
	"sync"
	"testing"
)

func TestSnowConsensus(t *testing.T) {
	totalNodes := 10
	nodes := make([]*node.Node, totalNodes)

	// Step 1: Initialize nodes
	for i := 0; i < totalNodes; i++ {
		nodes[i] = node.NewNode(i, totalNodes)
	}

	// Step 2: Start all nodes in parallel
	var wg sync.WaitGroup
	for _, n := range nodes {
		wg.Add(1)
		go func(node *node.Node) {
			defer wg.Done()
			log.Printf("Starting node %d", node.Id)
			node.Start()
			log.Printf("Node %d finished with preference: %s", node.Id, node.Consensus.GetPreference())
		}(n)
	}

	// Step 3: Wait for all nodes to finish
	wg.Wait()

	// Step 4: Validate that all nodes reach the same consensus
	preference := nodes[0].Consensus.GetPreference()
	for _, n := range nodes {
		if n.Consensus.GetPreference() != preference {
			t.Fatalf("Node %d did not reach consensus. Got %s", n.Id, n.Consensus.GetPreference())
		}
	}
	log.Printf("All nodes reached consensus on: %s", preference)
}
