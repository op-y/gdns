package web

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"

	"github.com/op-y/gdns/config"
	"github.com/op-y/gdns/web/api"
)

type Manager struct {
	server *http.Server
}

func NewManager(cfg *config.WebConfig) *Manager {

	r := gin.Default()
	r.GET("/ping", api.Pong)

	aGroup := r.Group("/a")
	aGroup.GET("/:domain", api.QueryA)
	aGroup.POST("/:domain", api.CreateA)
	aGroup.PUT("/:domain", api.UpdateA)
	aGroup.DELETE("/:domain", api.DeleteA)

	pprof.Register(r)

	mgr := &Manager{}
	srv := &http.Server{
		Addr:    cfg.Address,
		Handler: r,
	}
	mgr.server = srv
	return mgr
}

func (mgr *Manager) Start() {
	if err := mgr.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}

func (mgr *Manager) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := mgr.server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
}
