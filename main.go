// Simple example embedding and extending Gate.
package main

import (
	"context"
	"time"

	"github.com/robinbraemer/event"
	. "go.minekube.com/common/minecraft/component"
	"go.minekube.com/common/minecraft/component/codec/legacy"
	"go.minekube.com/gate/cmd/gate"
	"go.minekube.com/gate/pkg/edition/java/proxy"
)

func main() {
	// Add our "plug-in" to be initialized on Gate start.
	proxy.Plugins = append(proxy.Plugins, proxy.Plugin{
		Name: "SimpleProxy",
		Init: func(ctx context.Context, proxy *proxy.Proxy) error {
			return newSimpleProxy(proxy).init()
		},
	})

	// Execute Gate entrypoint and block until shutdown.
	// We could also run gate.Start if we don't need Gate's command-line.
	gate.Execute()
}

// SimpleProxy is a simple proxy to showcase some features of Gate.
//
// In this example:
//   - Add a `/broadcast` command
//   - Send a message when player switches the server
//   - Show boss bars to players
type SimpleProxy struct {
	*proxy.Proxy
}

var legacyCodec = &legacy.Legacy{Char: legacy.AmpersandChar}

func newSimpleProxy(proxy *proxy.Proxy) *SimpleProxy {
	return &SimpleProxy{
		Proxy: proxy,
	}
}

// initialize our sample proxy
func (p *SimpleProxy) init() error {
	p.registerSubscribers()
	return nil
}

// Register event subscribers
func (p *SimpleProxy) registerSubscribers() {
	// Change the MOTD response.
	event.Subscribe(p.Event(), 0, pingHandler())

}

func pingHandler() func(p *proxy.PingEvent) {
	motd := &Text{Content: "Simple Proxy!\nJoin and test me."}
	return func(e *proxy.PingEvent) {
		p := e.Ping()
		p.Description = motd
		p.Players.Max = p.Players.Online + 1
	}
}

// tick runs a function every interval until the context is cancelled.
func tick(ctx context.Context, interval time.Duration, fn func()) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			fn()
		case <-ctx.Done():
			return
		}
	}
}
