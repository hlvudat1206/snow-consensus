package tests

import (
	"snow-consensus/node"
	"testing"
)

func TestSnowConsensus(t *testing.T) {
	totalNodes := 10
	nodes := make([]*node.Node, totalNodes)

	for i := 0; i < totalNodes; i++ {
		nodes[i] = node.NewNode(i, totalNodes)
	}

	for _, n := range nodes {
		go n.Start()
	}

	// Check consensus (all nodes must agree)
	preference := nodes[0].Consensus.GetPreference()
	for _, n := range nodes {
		if n.Consensus.GetPreference() != preference {
			t.Fatalf("Node %d did not reach consensus. Got %s", n.ID, n.Consensus.GetPreference())
		}
	}
}
