package config

import (
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Storage *StorageConfig
	Cache   *CacheConfig
	DNS     *DNSConfig
	Web     *WebConfig
}

type StorageConfig struct {
	Type   string
	MySQL  *MysqlConfig
	SQLite *SQLiteConfig
}

type MysqlConfig struct {
	Host            string
	Port            int
	DB              string
	Username        string
	Password        string
	MaxOpenConn     int
	MaxIdleConn     int
	ConnMaxLifetime int
}

type SQLiteConfig struct {
	File string
}

type CacheConfig struct {
	Type  string
	Redis *RedisConfig
	Mem   *MemConfig
}

type RedisConfig struct {
	Host        string
	Port        int
	MaxActive   int
	MaxIdle     int
	IdleTimeout int
}

type MemConfig struct {
	InitSize int
}

type DNSConfig struct {
	Address    string
	Nameserver []string
	Timeout    int
	TTL        int
}

type WebConfig struct {
	Address string
}

var (
	DefaultConfig = &Config{
		Storage: DefaultStorageConfig,
		Cache:   DefaultCacheConfig,
		DNS:     DefaultDNSConfig,
		Web:     DefaultWebConfig,
	}

	DefaultStorageConfig = &StorageConfig{
		Type: "sqlite",
		MySQL: &MysqlConfig{
			MaxOpenConn:     100,
			MaxIdleConn:     50,
			ConnMaxLifetime: 3600,
		},
		SQLite: &SQLiteConfig{
			File: "test.db",
		},
	}

	DefaultCacheConfig = &CacheConfig{
		Type: "mem",
		Redis: &RedisConfig{
			MaxActive:   100,
			MaxIdle:     50,
			IdleTimeout: 30,
		},
		Mem: &MemConfig{
			InitSize: 1024,
		},
	}

	DefaultDNSConfig = &DNSConfig{
		Address: "127.0.0.1:53",
		Timeout: 10,
		TTL:     3600,
	}

	DefaultWebConfig = &WebConfig{
		Address: "127.0.0.1:8080",
	}
)

func Load(s string) (*Config, error) {
	cfg := &Config{}
	cfg = DefaultConfig

	_, err := toml.Decode(s, &cfg)
	return cfg, err
}

func LoadFile(fp string) *Config {
	fp = joinPath(fp)
	content, err := os.ReadFile(fp)
	if err != nil {
		panic(err)
	}
	cfg, err := Load(string(content))
	if err != nil {
		panic(err)
	}
	return cfg
}

func joinPath(fp string) string {
	if filepath.IsAbs(fp) {
		return fp
	}
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	return filepath.Join(dir, fp)
}
