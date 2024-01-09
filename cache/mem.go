package cache

import (
	"errors"
	"sync"

	"github.com/op-y/gdns/config"
)

const (
	MEM = "mem"
)

var (
	MemCacheMissErr = errors.New("domain not found in mem cache")
)

type mem struct {
	sync.RWMutex
	db map[string][]string
}

func NewMemCache(cfg *config.CacheConfig) Cache {
	db := make(map[string][]string, cfg.Mem.InitSize)
	m := &mem{
		db: db,
	}
	return m
}

func (m *mem) Name() string {
	return "mem"
}

func (m *mem) QueryA(domain string) ([]string, error) {
	m.RLock()
	defer m.RUnlock()
	if records, ok := m.db[domain]; ok {
		return records, nil
	}
	return nil, MemCacheMissErr
}

func (m *mem) UpdateA(domain string, records []string) error {
	m.Lock()
	defer m.Unlock()
	m.db[domain] = records
	return nil
}

func (m *mem) DeleteA(domain string) error {
	m.Lock()
	defer m.Unlock()
	delete(m.db, domain)
	return nil
}
