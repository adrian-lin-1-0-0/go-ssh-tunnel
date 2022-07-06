package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/adrian-lin-1-0-0/go-ssh-tunnel/pkg/tunnel"
)

var usageStr = `
Usage: ssh tunnel [options]
Common Options:
    -h, --help                       Show this message
    -u,                              ssh username
    -p,                              ssh password
    -i,                              identity file [Absolute Path]
    -s,                              ssh server address (127.0.0.1[:22] )
    -t,                              target address (:80)
    -l,                              local address (:80)
`

var (
	user    string
	pwd     string
	keyPath string
	svrAddr string
	srcAddr string
	dstAddr string
)

func init() {
	flag.StringVar(&user, "u", "", "ssh username")
	flag.StringVar(&pwd, "p", "", "ssh password")
	flag.StringVar(&keyPath, "i", "", "identity file")
	flag.StringVar(&svrAddr, "s", "127.0.0.1:22", "ssh server address")
	flag.StringVar(&dstAddr, "t", "0.0.0.0:80", "target address")
	flag.StringVar(&srcAddr, "l", "0.0.0.0:80", "local address")

	flag.Usage = usage
	flag.Parse()
	if user == "" {
		usage()
		os.Exit(0)
	}
}

func usage() {
	fmt.Printf("%s\n", usageStr)
}

func main() {
	if -1 == strings.Index(svrAddr, ":") {
		svrAddr += ":22"
	}
	tunnel.NewTunnel(user, pwd, keyPath, svrAddr, srcAddr, dstAddr)
}
