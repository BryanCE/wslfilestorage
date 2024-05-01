package main

import (
	"log"

	"github.com/BryanCE/firstmodule/filestorage/p2p"
)

func main ()  {
	tr := p2p.NewTcpTransport(":3000")
	if err := tr.ListenAndAccept(); err != nil {
		log.Fatalf("Error listening and accepting connections: %v", err)
	}

	select{}
}