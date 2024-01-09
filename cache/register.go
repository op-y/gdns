package cache

import (
	"github.com/op-y/gdns/config"
)

var (
	registry     = make(map[string]CacheBuilder)
	defaultCache Cache
)

type CacheBuilder func(*config.CacheConfig) Cache

func Register(name string, builder CacheBuilder) {
	registry[name] = builder
}

func NewCache(cfg *config.CacheConfig) Cache {
	cachebuilder := registry[cfg.Type]
	return cachebuilder(cfg)
}

func NewDefaultCache(cfg *config.CacheConfig) Cache {
	defaultCache = NewCache(cfg)
	return defaultCache
}

func init() {
	Register(MEM, NewMemCache)
}

func Default() Cache {
	return defaultCache
}
