// Represents a node in the network.
// Handles communication and consensus logic.

package node

import (
	"log"
	"math/rand"
	"snow-consensus/snow-consensus/consensus"
	"snow-consensus/snow-consensus/p2p"
	"sync"
	// "snow-consensus/consensus"
	// "snow-consensus/p2p"
)

type Node struct {
	Id        int
	network   *p2p.Network
	Consensus *consensus.Snow  // Expose the Snow consensus instance
	peers     []int            // List of peer IDs that this node communicates with
	inbox     chan interface{} // Channel for receiving messages
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
		Id:        id,
		network:   net,
		Consensus: consensus.NewSnow(randomPreference()),
		peers:     peers,
		inbox:     inbox,
	}
}

func (n *Node) Start() {
	for !n.Consensus.IsAccepted() {
		sample := n.samplePeers()                   // Randomly select peers
		preferences := n.collectPreferences(sample) // Gather preferences
		n.Consensus.Sample(preferences)
		// time.Sleep(10 * time.Millisecond) // Prevent busy-waiting
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
	var wg sync.WaitGroup
	mu := sync.Mutex{}

	// Collect preferences from peers concurrently
	for i, peer := range peers {
		wg.Add(1)
		go func(i, peer int) {
			defer wg.Done()

			// Send message to the peer
			n.network.SendMessage(peer, n.Consensus.GetPreference())

			// Read response from the inbox (non-blocking approach)
			select {
			case msg := <-n.inbox:
				if preference, ok := msg.(string); ok {
					mu.Lock()
					preferences[i] = preference
					mu.Unlock()
				} else {
					mu.Lock()
					preferences[i] = "" // Fallback for invalid message type
					mu.Unlock()
				}
			default:
				mu.Lock()
				preferences[i] = "" // Fallback for no response
				mu.Unlock()
			}
		}(i, peer)
	}

	// Wait for all goroutines to finish
	wg.Wait()
	return preferences
}

func randomPreference() string {
	if rand.Intn(2) == 0 {
		return "A"
	}
	return "B"
}
