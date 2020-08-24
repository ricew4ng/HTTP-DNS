package example

import (
	"context"
	"github.com/coredns/coredns/plugin/pkg/log"
	"github.com/miekg/dns"
)

type Hook struct {

}

func(h Hook)  ServeDNS(ctx context.Context, w dns.ResponseWriter, r *dns.Msg) (int, error) {
	log.Info("request received: %v", r)

	return 0, nil
}

func (h Hook) Name() string { return "hook" }