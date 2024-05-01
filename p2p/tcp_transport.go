package p2p

import (
	"fmt"
	"net"
	"sync"
)

// TCPPeer represents remote node
type TCPPeer struct {
	//conn is underlying TCP connection
	conn net.Conn

	//outbound is true if we dialed the connection, false if the connection was accepted
	outbound bool
}

type TcpTransport struct {
	listenAddress string
	listener 	net.Listener
	shakeHands HandshakeFunc
	decoder Decoder
	mu 			sync.RWMutex
	peer 		map[net.Addr]Peer
}

func NewTcpPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn: conn,
		outbound: outbound,
	}
}


func NewTcpTransport(listenAddr string) *TcpTransport {
	return &TcpTransport{
		shakeHands: NopHandShakeFunc,
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
		fmt.Printf("Accepted connection from %v\n", conn.RemoteAddr())
		
		go t.handleConn(conn)
	}
}

type Temp struct {}

func (t *TcpTransport) handleConn(conn net.Conn) {	
	peer := NewTcpPeer(conn, true)

	if err := t.shakeHands(conn); err != nil {
		fmt.Printf("Handshake failed with %v: %v\n", peer, err)
		return
	}


	//read loop for connection
	msg := &Temp{}
	for {
		if err := t.decoder.Decode(conn, msg); err != nil {
			fmt.Printf("Error decoding message from %v: %v\n", msg, err)
			continue
		}
	}
}