package main

import (
	"log"

	"github.com/BryanCE/firstmodule/filestorage/p2p"
)

func main ()  {
	tcpOpts := p2p.TcpTransportOptions{
		ListenAddr: ":4000",
		HandshakeFunc: p2p.NopHandShakeFunc,
		Decoder: p2p.GOBDecoder{},
	}
	tr := p2p.NewTcpTransport(tcpOpts)
	if err := tr.ListenAndAccept(); err != nil {
		log.Fatalf("Error listening and accepting connections: %v", err)
	}
}