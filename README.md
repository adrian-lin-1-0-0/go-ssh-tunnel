# SSH Tunnel


## Cmd 

### Quick Start

Install

```
go install -v github.com/adrian-lin-1-0-0/go-ssh-tunnel/cmd/sshtnl@latest
```

Usage

```
Usage: ssh tunnel [options]
Common Options:
    -h, --help                       Show this message
    -u,                              ssh username
    -p,                              ssh password
    -i,                              identity file [Absolute Path]
    -s,                              ssh server address (127.0.0.1[:22] )
    -t,                              target address (:80)
    -l,                              local address (:80)
```

e.g.

with identity file
```
sshtnl -u adrian -i /home/adrian/key/key.pem -s 10.0.0.1 -t :3306 -l :3306
```

with password
```
sshtnl -u adrian -p adrian-pwd -s 10.0.0.1 -t :3306 -l :3306
```

## Module

### Quick start

```go
package main

import (
	"github.com/adrian-lin-1-0-0/go-ssh-tunnel/pkg/tunnel"
)

func main() {
	user := "adrian"
	pwd:= "adrian-dev"
	keyPath := ""
	svrAddr := "10.0.0.1:22"
	srcAddr := ":3306"
	dstAddr := ":3306"
	tunnel.NewTunnel(user, pwd, "", svrAddr, srcAddr, dstAddr)
}
```
