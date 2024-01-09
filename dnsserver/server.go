package dnsserver

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/miekg/dns"

	"github.com/op-y/gdns/config"
)

type Manager struct {
	server *dns.Server
}

func NewManager(cfg *config.DNSConfig) *Manager {
	mgr := &Manager{}

	udpAddr, err := net.ResolveUDPAddr("udp", cfg.Address)
	if err != nil {
		panic(err)
	}
	p, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		panic(err)
	}
	tcpAddr, err := net.ResolveTCPAddr("tcp", cfg.Address)
	if err != nil {
		panic(err)
	}
	l, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		panic(err)
	}

	nameservers := make([]string, len(cfg.Nameserver))
	copy(nameservers, cfg.Nameserver)
	handler := &DefaultHandler{
		nameservers: nameservers,
		ttl:         cfg.TTL,
	}

	srv := &dns.Server{Listener: l, PacketConn: p, Handler: handler}
	mgr.server = srv

	return mgr
}

func (mgr *Manager) Start() {
	err := mgr.server.ActivateAndServe()
	if err != nil {
		panic(err)
	}
}

func (mgr *Manager) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := mgr.server.ShutdownContext(ctx); err != nil {
		log.Fatal("DNS Server Shutdown Error:", err)
	}
}
