package node

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"net"
)

type NodeID [32]byte

type Node struct {
	// user is set as a precautionary measure.
	// The same user can spin multiple nodes from the same machine in the same network.
	// Why would they want to participate with multiple nodes, beats me, but it may happen :)
	user string

	// ip is the parsed ip of a node.
	// Used to store ip of other nodes.
	// Pretty much useless when setting a node, apart from when using proxies.
	ip net.IP

	// port is the port on which a node listens on.
	port uint16

	// id is computed from the bytes of User, IP and Port of a node.
	// It uses SHA256.
	id NodeID
}

func (n *Node) User() string {
	return n.user
}

func (n *Node) Ip() net.IP {
	return n.ip
}

func (n *Node) Port() uint16 {
	return n.port
}

func (n *Node) Id() NodeID {
	return n.id
}

func (n *Node) String() string {
	return fmt.Sprintf("node: { user:%s ; ip:%s ; port:%d ; id:%x }", n.user, n.ip, n.port, n.id)
}

// NewNode creates a node instance to keep in memory.
// This function is used for storing incoming nodes' data.
// It should not be used for the locally run node.
func NewNode(user string, ip string, port uint16) (Node, error) {
	parsedIp := net.ParseIP(ip)
	if parsedIp == nil {
		return Node{}, fmt.Errorf("cannot create new node: ip is invalid: %s", ip)
	}
	buf := make([]byte, 0, len(parsedIp)+2+len(user))
	buf = append(buf, parsedIp...)
	buf = binary.BigEndian.AppendUint16(buf, port)
	buf = append(buf, []byte(user)...)

	return Node{user: user, ip: parsedIp, port: port, id: sha256.Sum256(buf)}, nil
}
