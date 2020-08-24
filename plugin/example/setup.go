package example

import (
	"github.com/caddyserver/caddy"
	"github.com/coredns/coredns/core/dnsserver"
	"github.com/coredns/coredns/plugin"
)

func init() { plugin.Register("example", setup) }

func setup(c *caddy.Controller) error {
	c.Next() // 'example'
	if c.NextArg() {
		return plugin.Error("example", c.ArgErr())
	}

	dnsserver.GetConfig(c).AddPlugin(func(next plugin.Handler) plugin.Handler {
		return Example{}
	})

	return nil
}