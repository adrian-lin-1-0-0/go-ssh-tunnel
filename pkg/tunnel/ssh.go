package tunnel

import (
	"errors"
	"io/ioutil"
	"net"
	"strings"
	"time"

	"golang.org/x/crypto/ssh"
)

func NewSSHConfig(user, pwd, keyPath string) (*ssh.ClientConfig, error) {
	var auth []ssh.AuthMethod

	if pwd != "" {
		auth = append(auth, ssh.Password(pwd))
	}

	if keyPath != "" {
		// private key, the crypto/x509 package
		key, err := ioutil.ReadFile(keyPath)
		if err != nil {
			return nil, errors.New(err.Error() + ",ReadFile :" + keyPath)
		}

		signer, err := ssh.ParsePrivateKey(key)
		if err != nil {
			return nil, err
		}
		auth = append(auth, ssh.PublicKeys(signer))
	}

	return &ssh.ClientConfig{
		User: user,
		Auth: auth,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
		Timeout: 30 * time.Second,
	}, nil
}

func NewSSHClient(config *ssh.ClientConfig, addr net.Addr) (*ssh.Client, error) {
	client, err := ssh.Dial("tcp", addr.String(), config)
	if err != nil {
		return nil, errors.New(err.Error() + ", ssh.dial addr " + addr.String())
	}
	return client, nil
}

func NewSSHConn(client *ssh.Client, addr net.Addr) (*net.Conn, error) {
	addrStr := addr.String()
	if strings.Index(addrStr, ":") == 0 {
		addrStr = "0.0.0.0" + addrStr
	}

	conn, err := client.Dial("tcp", addrStr)
	if err != nil {
		return nil, errors.New(err.Error() + ", ssh.Client : tcp.dial addr " + addr.String())
	}
	return &conn, nil
}
