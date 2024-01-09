package storage

import (
	"github.com/op-y/gdns/config"
)

var (
	registry       = make(map[string]StorageBuilder)
	defaultStorage Storage
)

type StorageBuilder func(*config.StorageConfig) Storage

func Register(name string, builder StorageBuilder) {
	registry[name] = builder
}

func NewStorage(cfg *config.StorageConfig) Storage {
	cachebuilder := registry[cfg.Type]
	return cachebuilder(cfg)
}

func NewDefaultStorage(cfg *config.StorageConfig) Storage {
	defaultStorage = NewStorage(cfg)
	return defaultStorage
}

func init() {
	Register(SQLITE, NewSqlletStorage)
}

func Default() Storage {
	return defaultStorage
}
