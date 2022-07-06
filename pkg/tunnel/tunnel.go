package tunnel

import (
	"io"
	"log"
	"net"

	"golang.org/x/crypto/ssh"
)

type Tunnel struct {
	SSHConfig *ssh.ClientConfig
	SvrAddr   net.Addr
	SrcAddr   net.Addr
	DstAddr   net.Addr
}

func NewTunnel(user, pwd, keyPath, svrAddr, srcAddr, dstAddr string) {
	if svrAddr == "" {
		svrAddr = ":22"
	}
	config := &Config{
		User:    user,
		Pwd:     pwd,
		KeyPath: keyPath,
		SvrAddr: svrAddr,
		SrcAddr: srcAddr,
		DstAddr: dstAddr,
	}
	t, err := New(*config)
	if err != nil {
		log.Fatalln("Create tunnel failed ", err.Error())
	}
	if err = t.Run(); err != nil {
		log.Fatalln("Run tunnel failed ", err.Error())
	}
}

func New(config Config) (*Tunnel, error) {
	sshConfig, err := NewSSHConfig(config.User, config.Pwd, config.KeyPath)
	if err != nil {
		return nil, err
	}

	svrAddr, err := net.ResolveTCPAddr("tcp", config.SvrAddr)
	if err != nil {
		return nil, err
	}

	srcAddr, err := net.ResolveTCPAddr("tcp", config.SrcAddr)
	if err != nil {
		return nil, err
	}

	dstAddr, err := net.ResolveTCPAddr("tcp", config.DstAddr)
	if err != nil {
		return nil, err
	}

	return &Tunnel{
		SSHConfig: sshConfig,
		SvrAddr:   svrAddr,
		SrcAddr:   srcAddr,
		DstAddr:   dstAddr,
	}, nil

}

func (t *Tunnel) Run() error {

	sshClinet, err := NewSSHClient(t.SSHConfig, t.SvrAddr)
	if err != nil {
		return err
	}

	tcpListener, err := NewTCPListener(t.SrcAddr)
	if err != nil {
		return err
	}
	defer (*tcpListener).Close()

	for {
		tcpConn, err := (*tcpListener).Accept()
		if err != nil {
			if opErr, ok := err.(*net.OpError); ok && opErr.Timeout() {
				continue
			}
			return err
		}
		sshConn, err := NewSSHConn(sshClinet, t.DstAddr)
		if err != nil {
			return err
		}
		defer (*sshConn).Close()
		log.Printf("%s -> (%s)%s", tcpConn.RemoteAddr().String(), t.SvrAddr.String(), t.DstAddr.String())
		go pipe(tcpConn, *sshConn)
	}
}

func pipe(src, dst net.Conn) {
	defer src.Close()

	go func() {
		defer dst.Close()
		io.Copy(dst, src)
	}()

	io.Copy(src, dst)
}
