package tunnel

import "net"

func NewTCPListener(add net.Addr) (*net.Listener, error) {
	listen, err := net.Listen("tcp", add.String())
	if err != nil {
		return nil, err
	}
	return &listen, nil
}
