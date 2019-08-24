package Network

import "net"

type NetConnection struct {
	connectionID int
	conn         *net.Conn
}
