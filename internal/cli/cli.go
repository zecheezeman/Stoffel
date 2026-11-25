// SPDX-License-Identifier: MPL-2.0
package cli

import (
	"flag"
	"fmt"
	"math"
)

type Opts struct {
	// ConnectionIp indicates the IP we are connecting to when we first connect to a network.
	ConnectionIp string
	// ConnectionPort indicates the port we are connecting to when we first connect to a network.
	ConnectionPort int
	// NewNetwork indicates that the we are participating to a new network.
	NewNetwork bool
	// PrimaryConnectionCap indicates the maximum number of primary connections a node can have.
	PrimaryConnectionCap int
	// ListenPort
	ListenPort int
}

const (
	defaultString = ""
	defaultInt    = -1
	defaultBool   = false

	maxPrimaryConnection = 255
	defaultListenPort    = 8080
)

func ReadOpts() (Opts, error) {
	o := Opts{}

	defaultOpts := Opts{
		ConnectionIp:         defaultString,
		ConnectionPort:       defaultInt,
		NewNetwork:           defaultBool,
		PrimaryConnectionCap: defaultInt,
		ListenPort:           defaultListenPort,
	}

	flag.StringVar(&o.ConnectionIp, "cip", defaultString, "the IP we are connecting to when we first connect to a network")
	flag.IntVar(&o.ConnectionPort, "cport", defaultInt, "the port we are connecting to when we first connect to a network")
	flag.BoolVar(&o.NewNetwork, "newnet", defaultBool, "indicates that the we are participating to a new network")
	flag.IntVar(&o.PrimaryConnectionCap, "primcap", defaultInt, "the maximum number of primary connections a node can have - max 255 primary connections")
	flag.IntVar(&o.ListenPort, "lport", defaultListenPort, "the port we are listening on locally")

	flag.Parse()

	if o == defaultOpts {
		return o, fmt.Errorf("no configuration provided")
	}

	if err := o.validate(); err != nil {
		return o, fmt.Errorf("invalid configuration: %s", err.Error())
	}
	return o, nil
}

func (o *Opts) validate() error {
	if (o.ConnectionIp == defaultString || o.ConnectionPort == defaultInt) && o.NewNetwork {
		return fmt.Errorf("new network creation requested but connection information is invalid: ip:%s port:%d", o.ConnectionIp, o.ConnectionPort)
	}

	if (o.ConnectionIp != defaultString || o.ConnectionPort != defaultInt) && !o.NewNetwork {
		return fmt.Errorf("connection information provided but no new network request: ip:%s port:%d", o.ConnectionIp, o.ConnectionPort)
	}

	if o.PrimaryConnectionCap == defaultInt {
		return fmt.Errorf("primary connection capacity must be specified at the start of each node")
	}

	if o.PrimaryConnectionCap > maxPrimaryConnection {
		return fmt.Errorf("a node can store at most %d primary connections. requested: %d", maxPrimaryConnection, o.PrimaryConnectionCap)
	} else if o.PrimaryConnectionCap < 0 {
		return fmt.Errorf("negative \"primcap\" value passed: %d", o.PrimaryConnectionCap)
	}

	if o.ListenPort > math.MaxUint16 || o.ListenPort < 0 {
		return fmt.Errorf("invalid \"lport\" value: %d", o.ListenPort)
	}

	return nil
}
