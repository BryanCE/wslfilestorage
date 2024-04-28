package p2p

import (
	"net"
	"sync"
)

type TcpTransport struct {
	listenAddress string
	listener 	net.Listener
	mu 			sync.RWMutex
	peer 		map[net.Addr]Peer
}

func NewTcpTransport(listenAddr string) *TcpTransport {
	return &TcpTransport{
		listenAddress: listenAddr,
		peer: make(map[net.Addr]Peer),
	}
}