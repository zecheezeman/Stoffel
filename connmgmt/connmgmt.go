package connmgmt

import (
	"fmt"

	"github.com/zecheezeman/Stoffel/node"
)

// ConnectionManager is supposed to store all conections and views of the current node.
// Must use `New` method to instantiate this struct.
// The behaviour should be this: NodeID -> slice of Node, where the first element is the actual element.
// Ex: conns[nid][0] == Node, where nid == Node.Id()
// TODO: Find a smarter way to store connections, possibly without data duplication
type ConnectionManager struct {
	conns map[node.NodeID][]node.Node
}

// New instantiates a new ConnectionManager with no previously allocalted capacity.
func New() ConnectionManager {
	return ConnectionManager{
		// Leave without size as of now
		conns: make(map[node.NodeID][]node.Node),
	}
}

// AddPrimary stores a new primary connection.
func (cm *ConnectionManager) AddPrimary(n node.Node) error {
	if _, ok := cm.conns[n.Id()]; ok {
		return fmt.Errorf("node id already exists in connection manager: %x", n.Id())
	}
	cm.conns[n.Id()] = make([]node.Node, 0)
	cm.conns[n.Id()] = append(cm.conns[n.Id()], n)
	return nil
}

// AddNonPrimary stores a new connection that is not a primary connection.
func (cm *ConnectionManager) AddNonPrimary(primId node.NodeID, n node.Node) error {
	if _, ok := cm.conns[primId]; !ok {
		return fmt.Errorf("primary connection node id does not exist: %x", n.Id())
	}

	if _, ok := cm.conns[n.Id()]; ok {
		return fmt.Errorf("non-primary node id already exists in connection manager: %x", n.Id())
	}

	cm.conns[primId] = append(cm.conns[primId], n)
	cm.conns[n.Id()] = make([]node.Node, 0)
	cm.conns[n.Id()] = append(cm.conns[n.Id()], n)

	return nil
}

func (cm *ConnectionManager) Get(nid node.NodeID) (node.Node, error) {
	if n, ok := cm.conns[nid]; ok {
		return n[0], nil
	}
	return node.Node{}, fmt.Errorf("could not find node with id: %x", nid)
}
