package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	flag "github.com/spf13/pflag"

	"github.com/op-y/gdns/cache"
	"github.com/op-y/gdns/config"
	"github.com/op-y/gdns/dnsserver"
	"github.com/op-y/gdns/storage"
	"github.com/op-y/gdns/web"
)

var (
	cfg        *config.Config
	dnscache   cache.Cache
	dnsstorage storage.Storage
	dnsManager *dnsserver.Manager
	webManager *web.Manager
)

func main() {
	var fp string
	flag.StringVarP(&fp, "conf", "f", "dns.toml", "configuration file")
	flag.Lookup("conf").NoOptDefVal = "dns.toml"
	flag.Parse()

	cfg = config.LoadFile(fp)

	dnscache = cache.NewDefaultCache(cfg.Cache)
	dnsstorage = storage.NewDefaultStorage(cfg.Storage)

	webManager = web.NewManager(cfg.Web)
	dnsManager = dnsserver.NewManager(cfg.DNS)

	go dnsManager.Start()
	go webManager.Start()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	select {
	case <-quit:
		dnsManager.Stop()
		webManager.Stop()
		log.Println("gdns exiting")
	}
}
