package p2p

import (
	"fmt"
	"net"
	"sync"
)

// TcpPeer represents remote node
type TcpPeer struct {
	//conn is underlying TCP connection
	conn net.Conn

	//outbound is true if we dialed the connection, false if the connection was accepted
	outbound bool
}

type TcpTransportOptions struct {
	ListenAddr string
	HandshakeFunc HandshakeFunc
	Decoder Decoder
}

type TcpTransport struct {
	TcpTransportOptions
	listener 			net.Listener
	mu 					sync.RWMutex
	peers 				map[net.Addr]Peer
}

func NewTcpPeer(conn net.Conn, outbound bool) *TcpPeer {
	return &TcpPeer{
		conn: conn,
		outbound: outbound,
	}
}


func NewTcpTransport(opts TcpTransportOptions) *TcpTransport {
	return &TcpTransport{
		TcpTransportOptions: opts,
	}
}

func (t *TcpTransport) ListenAndAccept() error {
	var err error
	t.listener, err = net.Listen("tcp", t.ListenAddr)
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

	if err := t.HandshakeFunc(peer); err != nil {
		fmt.Printf("Handshake failed with %v: %v\n", peer, err)
		return 
	}


	//read loop for connection
	// msg := &Temp{}

	buf := make([]byte, 2000)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Printf("TCP error: %s\n", err)
		}
		// if err := t.Decoder.Decode(conn, msg); err != nil {
		// 	fmt.Printf("Error decoding message from %v: %v\n", msg, err)
		// 	continue
		// }	
		fmt.Printf("Message: %v\n", buf[:n])	
	}
}