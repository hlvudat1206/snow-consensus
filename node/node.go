package node

import (
	"math/rand"
	"snow-consensus/snow-consensus/consensus"
	"snow-consensus/snow-consensus/p2p"
	// "snow-consensus/consensus"
	// "snow-consensus/p2p"
)

type Node struct {
	id        int
	network   *p2p.Network
	consensus *consensus.Snow
	peers     []int
	inbox     chan interface{}
}

func NewNode(id, totalNodes int) *Node {
	net := p2p.NewNetwork()
	inbox := net.RegisterNode(id)

	peers := make([]int, totalNodes-1)
	idx := 0
	for i := 0; i < totalNodes; i++ {
		if i != id {
			peers[idx] = i
			idx++
		}
	}

	return &Node{
		id:        id,
		network:   net,
		consensus: consensus.NewSnow(randomPreference()),
		peers:     peers,
		inbox:     inbox,
	}
}

func (n *Node) Start() {
	for !n.consensus.IsAccepted() {
		sample := n.samplePeers()
		preferences := n.collectPreferences(sample)
		n.consensus.Sample(preferences)
	}
}

func (n *Node) samplePeers() []int {
	sample := make([]int, consensus.SampleSize)
	for i := 0; i < consensus.SampleSize; i++ {
		sample[i] = n.peers[rand.Intn(len(n.peers))]
	}
	return sample
}

func (n *Node) collectPreferences(peers []int) []string {
	preferences := make([]string, len(peers))
	for i, peer := range peers {
		n.network.SendMessage(peer, n.consensus.GetPreference())
		preferences[i] = <- n.inbox
	}
	return preferences
}

func randomPreference() string {
	if rand.Intn(2) == 0 {
		return "A"
	}
	return "B"
}
