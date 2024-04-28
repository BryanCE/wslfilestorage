package p2p

import (
	"fmt"
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
	}
}

func (t *TcpTransport) ListenAndAccept() error {
	var err error
	t.listener, err = net.Listen("tcp", t.listenAddress)
	if err != nil {
		return err
	}
	go t.startAcceptLoop()
	return nil
}

func (t *TcpTransport) startAcceptLoop() {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Printf("TCP Error accepting connection: %v\n", err)
		}
		go t.handleConnn(conn)
	}
}

func (t *TcpTransport) handleConnn(conn net.Conn) {	
	fmt.Printf("New incoming connection from %v\n", conn)
}