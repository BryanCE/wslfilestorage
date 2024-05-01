package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTcpTransport(t *testing.T) {
	listenAddr := ":3000"
	tr := NewTcpTransport(listenAddr)
	assert.Equal(t, tr.listenAddress, listenAddr)

	assert.Nil(t, tr.ListenAndAccept())

	select {

	}
}