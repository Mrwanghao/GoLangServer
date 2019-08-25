package netserver

import "net"

type NetConnection struct {
	connectionID int
	conn         *net.Conn
}

func NewNetConnection(connectionID int, conn *net.Conn) *NetConnection {
	var netConnection *NetConnection = new(NetConnection)
	netConnection.connectionID = connectionID
	netConnection.conn = conn
	return netConnection
}
