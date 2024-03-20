package connection

import "net"

type GatewayConn struct {
	conn net.Conn
}

func (gc *GatewayConn) Init(c net.Conn) error {
	gc.conn = c
	return nil
}

func (gc *GatewayConn) Serve() {

}

func (gc *GatewayConn) Exit() {
	gc.conn.Close()
}
