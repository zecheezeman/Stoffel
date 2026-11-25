// SPDX-License-Identifier: MPL-2.0
package cli

import (
	"flag"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func readOptsWithArgs(args ...string) (Opts, error) {
	originalCommandLine := flag.CommandLine
	originalArgs := os.Args
	defer func() {
		flag.CommandLine = originalCommandLine
		os.Args = originalArgs
	}()

	flag.CommandLine = flag.NewFlagSet("stoffel", flag.ContinueOnError)
	flag.CommandLine.SetOutput(io.Discard)
	os.Args = append([]string{"stoffel"}, args...)

	return ReadOpts()
}

func TestReadOpts(t *testing.T) {
	opts, err := readOptsWithArgs("-primcap", "10")

	assert.NoError(t, err, "configuration should be read successfully")
	assert.Equal(t, Opts{
		ConnectionIp:         defaultString,
		ConnectionPort:       defaultInt,
		NewNetwork:           defaultBool,
		PrimaryConnectionCap: 10,
		ListenPort:           defaultListenPort,
	}, opts)
}

func TestReadOpts_AllOptions(t *testing.T) {
	opts, err := readOptsWithArgs(
		"-cip", "127.0.0.1",
		"-cport", "8081",
		"-newnet",
		"-primcap", "255",
		"-lport", "65535",
	)

	assert.NoError(t, err, "configuration should be read successfully")
	assert.Equal(t, Opts{
		ConnectionIp:         "127.0.0.1",
		ConnectionPort:       8081,
		NewNetwork:           true,
		PrimaryConnectionCap: maxPrimaryConnection,
		ListenPort:           65535,
	}, opts)
}

func TestReadOpts_NoConfiguration(t *testing.T) {
	_, err := readOptsWithArgs()

	assert.Error(t, err, "configuration should be required")
}

func TestReadOpts_NewNetworkWithoutConnectionInformation(t *testing.T) {
	_, err := readOptsWithArgs("-newnet", "-primcap", "10")

	assert.Error(t, err, "should be invalid configuration")
}

func TestReadOpts_ConnectionInformationWithoutNewNetwork(t *testing.T) {
	_, err := readOptsWithArgs("-cip", "127.0.0.1", "-cport", "8081", "-primcap", "10")

	assert.Error(t, err, "connection information should require a new network request")
}

func TestReadOpts_PrimaryConnectionCapacityNotSpecified(t *testing.T) {
	_, err := readOptsWithArgs("-lport", "8081")

	assert.Error(t, err, "primary connection capacity should be required")
}

func TestReadOpts_PrimaryConnectionCapacityTooHigh(t *testing.T) {
	_, err := readOptsWithArgs("-primcap", "256")

	assert.Error(t, err, "primary connection capacity should not exceed the maximum")
}

func TestReadOpts_NegativePrimaryConnectionCapacity(t *testing.T) {
	_, err := readOptsWithArgs("-primcap=-2")

	assert.Error(t, err, "primary connection capacity should not be negative")
}

func TestReadOpts_ListenPortTooHigh(t *testing.T) {
	_, err := readOptsWithArgs("-primcap", "10", "-lport", "65536")

	assert.Error(t, err, "listen port should not exceed the maximum")
}

func TestReadOpts_NegativeListenPort(t *testing.T) {
	_, err := readOptsWithArgs("-primcap", "10", "-lport=-1")

	assert.Error(t, err, "listen port should not be negative")
}
