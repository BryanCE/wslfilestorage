package p2p

// HandshakeFunc is a function that performs a handshake with a remote peer
type HandshakeFunc func(Peer) error

// NopHandShakeFunc is a no-op handshake function
func NopHandShakeFunc(Peer) error { return nil}