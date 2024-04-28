package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTcpTransport(t *testing.T) {
	listenAddr := "localhost:4000"
	tr := NewTcpTransport(listenAddr)
	assert.Equal(t, listenAddr, tr.listenAddress)
}