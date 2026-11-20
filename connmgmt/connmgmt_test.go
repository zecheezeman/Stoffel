// SPDX-License-Identifier: MPL-2.0
package connmgmt

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zecheezeman/Stoffel/node"
)

func TestAddPrimary(t *testing.T) {
	cm := New()

	n, err := node.NewNode("user1", "127.0.0.1", 8080)

	assert.NoError(t, err, "node should be created successfully")

	err = cm.AddPrimary(n)
	assert.NoError(t, err, "primary connection should be added successfully")
}

func TestAddPrimary_AlreadyExistent(t *testing.T) {
	cm := New()

	n, err := node.NewNode("user1", "127.0.0.1", 8080)

	assert.NoError(t, err, "node should be created successfully")

	err = cm.AddPrimary(n)
	assert.NoError(t, err, "primary connection should be added successfully")

	err = cm.AddPrimary(n)
	assert.Error(t, err, "primary connection should already exist")
}

func TestAddNonPrimary(t *testing.T) {
	cm := New()

	n, err := node.NewNode("user1", "127.0.0.1", 8080)

	assert.NoError(t, err, "node should be created successfully")

	err = cm.AddPrimary(n)
	assert.NoError(t, err, "primary connection should be added successfully")

	n1, err := node.NewNode("user2", "127.0.0.1", 8081)
	assert.NoError(t, err, "node should be created successfully")

	err = cm.AddNonPrimary(n.Id(), n1)
	assert.NoError(t, err, "non primary connection should be added successfully")
}

func TestAddNonPrimary_AlreadyExistent(t *testing.T) {
	cm := New()

	n, err := node.NewNode("user1", "127.0.0.1", 8080)

	assert.NoError(t, err, "node should be created successfully")

	err = cm.AddPrimary(n)
	assert.NoError(t, err, "primary connection should be added successfully")

	n1, err := node.NewNode("user2", "127.0.0.1", 8081)
	assert.NoError(t, err, "node should be created successfully")

	err = cm.AddNonPrimary(n.Id(), n1)
	assert.NoError(t, err, "non primary connection should be added successfully")

	err = cm.AddNonPrimary(n.Id(), n1)
	assert.Error(t, err, "non primary connection should already exist")
}

func TestAddNonPrimary_PrimaryConnectionDoesNotExist(t *testing.T) {
	cm := New()

	n, err := node.NewNode("user1", "127.0.0.1", 8080)
	assert.NoError(t, err, "node should be created successfully")

	n1, err := node.NewNode("user2", "127.0.0.1", 8081)
	assert.NoError(t, err, "node should be created successfully")

	err = cm.AddNonPrimary(n.Id(), n1)
	assert.Error(t, err, "primary connection should not exist")
}

func TestGet_PrimaryConnection(t *testing.T) {
	cm := New()

	n, err := node.NewNode("user1", "127.0.0.1", 8080)

	assert.NoError(t, err, "node should be created successfully")

	err = cm.AddPrimary(n)
	assert.NoError(t, err, "primary connection should be added successfully")

	ndup, err := cm.Get(n.Id())
	assert.NoError(t, err, "primary connection should exist")
	assert.Equal(t, n, ndup)
}

func TestGet_NonPrimaryConnection(t *testing.T) {
	cm := New()

	n, err := node.NewNode("user1", "127.0.0.1", 8080)

	assert.NoError(t, err, "node should be created successfully")

	err = cm.AddPrimary(n)
	assert.NoError(t, err, "primary connection should be added successfully")

	ndup, err := cm.Get(n.Id())
	assert.NoError(t, err, "primary connection should exist")
	assert.Equal(t, n, ndup)

	n1, err := node.NewNode("user2", "127.0.0.1", 8081)
	assert.NoError(t, err, "node should be created successfully")

	err = cm.AddNonPrimary(n.Id(), n1)
	assert.NoError(t, err, "primary connection should exist")

	n1dup, err := cm.Get(n1.Id())
	assert.NoError(t, err, "non primary connection should exist")
	assert.Equal(t, n1, n1dup)
}

func TestGet_DoesNotExist(t *testing.T) {
	cm := New()

	n, err := node.NewNode("user1", "127.0.0.1", 8080)

	assert.NoError(t, err, "node should be created successfully")

	_, err = cm.Get(n.Id())
	assert.Error(t, err, "primary connection should not exist")
}
