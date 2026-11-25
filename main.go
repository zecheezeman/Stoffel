// SPDX-License-Identifier: MPL-2.0
package main

import (
	"fmt"

	"github.com/zecheezeman/Stoffel/internal/cli"
	"github.com/zecheezeman/Stoffel/internal/node"
)

func main() {
	if opts, err := cli.ReadOpts(); err != nil {
		fmt.Printf("error while reading configuration options: %s\n", err)
		return
	} else {
		// TODO: Refactor this, it looks horrendous
		if err = node.RunLocalNode(opts); err != nil {
			fmt.Printf("error while running local node: %s\n", err)
		}
	}

}
