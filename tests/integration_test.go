package tests

import (
	// "snow-consensus/node"
	"snow-consensus/snow-consensus/node"
	"testing"
)

func TestSnowConsensus(t *testing.T) {
	totalNodes := 10
	nodes := make([]*node.Node, totalNodes)

	// Initialize nodes
	for i := 0; i < totalNodes; i++ {
		nodes[i] = node.NewNode(i, totalNodes)
	}

	// Start all nodes
	for _, n := range nodes {
		go n.Start()
	}

	// Check consensus
	preference := nodes[0].Consensus.GetPreference()
	for _, n := range nodes {
		if n.Consensus.GetPreference() != preference {
			t.Fatalf("Node %d did not reach consensus. Got %s", n.Id, n.Consensus.GetPreference())
		}
	}
}
