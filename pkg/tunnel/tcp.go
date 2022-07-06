package tunnel

import (
	"errors"
	"net"
	"strings"
)

func NewTCPListener(addr net.Addr) (*net.Listener, error) {
	listen, err := net.Listen("tcp", addr.String())
	if err != nil {
		return nil, err
	}
	return &listen, nil
}

func NewTcpConn(addr net.Addr) (*net.Conn, error) {
	addrStr := addr.String()
	if strings.Index(addrStr, ":") == 0 {
		addrStr = "0.0.0.0" + addrStr
	}

	conn, err := net.Dial("tcp", addrStr)
	if err != nil {
		return nil, errors.New(err.Error() + ", tcp.Client : tcp.dial addr " + addr.String())
	}
	return &conn, nil
}
