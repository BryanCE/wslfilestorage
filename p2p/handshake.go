package p2p

import "errors"

var ErrInvalidHandshake = errors.New("invalid handshake")

// HandshakeFunc is a function that performs a handshake with a remote peer
type HandshakeFunc func(Peer) error

// NopHandShakeFunc is a no-op handshake function
func NopHandShakeFunc(Peer) error { return nil}