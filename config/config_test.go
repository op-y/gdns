package config

import (
	"testing"
)

func TestLoadConfig(t *testing.T) {
	filename := "../config.toml"
	cfg := LoadFile(filename)
	t.Logf("%+v\n", cfg.Storage.MySQL)
	t.Logf("%+v\n", cfg.Storage.SQLite)
	t.Logf("%+v\n", cfg.Cache.Redis)
	t.Logf("%+v\n", cfg.Cache.Mem)
	t.Logf("%+v\n", cfg.DNS)
	t.Logf("%+v\n", cfg.Web)
}

func TestJoinPath(t *testing.T) {
	filename := "config.toml"
	fp := joinPath(filename)
	t.Logf(fp)
}
