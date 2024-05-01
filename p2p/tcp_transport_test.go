package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTcpTransport(t *testing.T) {
	opts := TcpTransportOptions{
		ListenAddr: ":4000",
		HandshakeFunc: NopHandShakeFunc,
		Decoder: GOBDecoder{},
	}
	listenAddr := ":4000"
	tr := NewTcpTransport(opts)
	assert.Equal(t, tr, listenAddr)

	assert.Nil(t, tr.ListenAndAccept())
}