// 1. This file implements the P2P networking layer.

// 2. Purpose
// + Provides a way for nodes to send and receive messages.
// + Simulates a broadcast system where nodes communicate through shared channels (transactions)

package p2p

import (
	"sync"
)

type Network struct {
	mu    sync.Mutex
	Nodes map[int]chan interface{}
}

func NewNetwork() *Network {
	return &Network{
		Nodes: make(map[int]chan interface{}),
	}
}

func (net *Network) RegisterNode(id int) chan interface{} {
	net.mu.Lock()
	defer net.mu.Unlock()
	channel := make(chan interface{}, 100)
	net.Nodes[id] = channel
	return channel
}

// Sends a message to a specific node
func (net *Network) SendMessage(to int, msg interface{}) {
	net.mu.Lock()
	defer net.mu.Unlock()
	if ch, exists := net.Nodes[to]; exists {
		ch <- msg
	}
}

// Sends a message to all nodes except the senders
func (net *Network) BroadcastMessage(from int, msg interface{}) {
	net.mu.Lock()
	defer net.mu.Unlock()
	for id, ch := range net.Nodes {
		if id != from {
			ch <- msg
		}
	}
}
