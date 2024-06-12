package p2p

// Peer is an interface that represents the remote node
type Peer interface {
	Close() error
}

// Transport is anything that handles tha communication
// between the nodes in the network. This can be of the
// form (TCP, USP, WebSockets, ...)
type Transport interface {
	ListenAndAccept() error
	Consume() <-chan RPC
}
