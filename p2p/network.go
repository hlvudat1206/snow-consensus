package p2p

import (
	"sync"
)

type Network struct {
	mu    sync.Mutex
	nodes map[int]chan interface{}
}

func NewNetwork() *Network {
	return &Network{
		nodes: make(map[int]chan interface{}),
	}
}

func (net *Network) RegisterNode(id int) chan interface{} {
	net.mu.Lock()
	defer net.mu.Unlock()
	channel := make(chan interface{}, 100)
	net.nodes[id] = channel
	return channel
}

func (net *Network) SendMessage(to int, msg interface{}) {
	net.mu.Lock()
	defer net.mu.Unlock()
	if ch, exists := net.nodes[to]; exists {
		ch <- msg
	}
}

func (net *Network) BroadcastMessage(from int, msg interface{}) {
	net.mu.Lock()
	defer net.mu.Unlock()
	for id, ch := range net.nodes {
		if id != from {
			ch <- msg
		}
	}
}
