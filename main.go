// Simple example embedding and extending Gate.
package main

import (
	"github.com/morinoparty/suspension_bridge/plugin/motd"
	"go.minekube.com/gate/cmd/gate"
	"go.minekube.com/gate/pkg/edition/java/proxy"
)

func main() {
	proxy.Plugins = append(proxy.Plugins, motd.Plugin)

	gate.Execute()
}
