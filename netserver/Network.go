package netserver

import (
	"fmt"
	"net"
)

type Network struct {
	conns map[int]*NetConnection
}

func test() {
	fmt.Println("你好")
}

var network *Network

func init() {
	network = new(Network)
}

func (network Network) AddNetConnection(conn *net.Conn) {
	var unUsedID int = 0
	for {

		for connectionID, _ := range network.conns {
			if connectionID == unUsedID {
				unUsedID += 1
				break
			}
		}

		var netConn *NetConnection = new(NetConnection)
		netConn.connectionID = unUsedID
		netConn.conn = conn
		// 能用，因为没有人使用
		network.conns[unUsedID] = netConn
		return
	}
}
