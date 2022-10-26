package network

import "bytes"


type NetHeader struct {

}
type RPC struct {
	From 	string
	Payload []byte
}

type Transports interface {
	Consume() <- chan RPC
	Connect(Transports) error
	SendMessage(NetHeader, []byte) error
	Addr () NetHeader
}