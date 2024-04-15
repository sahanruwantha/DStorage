package p2p

type HandshakeFunc func(Peer) error

func NOHandshakeFunc(any) error { return nil }
