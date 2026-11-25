// SPDX-License-Identifier: MPL-2.0
package connmgmt

import (
	"fmt"
	"slices"

	"github.com/zecheezeman/Stoffel/internal/node"
)

// Connection represents the single link between the primary connection of a node and the node itself.
type Connection struct {
	// In case the 'primary' node is nil, it means that the local node is the primary link of it.
	primary *node.Node
	node    node.Node
}

// ConnectionManager is supposed to store all conections and views of the current node.
// Must use `New` method to instantiate this struct.
// To find a node takes O(n), which is not the best way but it is the safest way as of now.
type ConnectionManager struct {
	conns []Connection
}

// New instantiates a new ConnectionManager with no previously allocalted capacity.
func New() ConnectionManager {
	return ConnectionManager{
		// Leave without size as of now
		conns: make([]Connection, 0),
	}
}

// AddPrimary stores a new primary connection.
func (cm *ConnectionManager) AddPrimary(n node.Node) error {
	if _, err := cm.Get(n.Id()); err == nil {
		return fmt.Errorf("node id already exists in connection manager: %x", n.Id())
	}

	cm.conns = append(cm.conns, Connection{
		primary: nil,
		node:    n,
	})

	return nil
}

// AddNonPrimary stores a new connection that is not a primary connection.
func (cm *ConnectionManager) AddNonPrimary(primId node.NodeID, n node.Node) error {
	var primConn node.Node
	var err error

	if primConn, err = cm.Get(primId); err != nil {
		return fmt.Errorf("primary connection node id does not exist: %x", n.Id())
	}

	if _, err := cm.Get(n.Id()); err == nil {
		return fmt.Errorf("non-primary node id already exists in connection manager: %x", n.Id())
	}

	cm.conns = append(cm.conns, Connection{
		primary: &primConn,
		node:    n,
	})

	return nil
}

func (cm *ConnectionManager) Get(nid node.NodeID) (node.Node, error) {
	if idx := slices.IndexFunc(
		cm.conns,
		func(c Connection) bool { return c.node.Id() == nid }); idx != -1 {
		return cm.conns[idx].node, nil
	}
	return node.Node{}, fmt.Errorf("no node found with id: %x", nid)
}
