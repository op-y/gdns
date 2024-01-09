package storage

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/op-y/gdns/config"
	"github.com/op-y/gdns/model"
)

var (
	SQLITE = "sqlite"
)

type sqlitedb struct {
	db *gorm.DB
}

func NewSqlletStorage(cfg *config.StorageConfig) Storage {
	s := &sqlitedb{}

	db, err := gorm.Open(sqlite.Open(cfg.SQLite.File), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	s.db = db

	if err := s.InitSQLiteSchema(); err != nil {
		panic(err)
	}

	return s
}

func (s *sqlitedb) InitSQLiteSchema() error {
	if !s.db.Migrator().HasTable(&model.A{}) {
		if err := s.db.Migrator().CreateTable(&model.A{}); err != nil {
			return err
		}
	}
	return nil
}

func (s *sqlitedb) Name() string {
	return "sqllet"
}

func (s *sqlitedb) QueryA(domain string) ([]string, error) {
	var As []model.A
	s.db.Select("record").Where("domain = ? AND record_type = ?", domain, "A").Find(&As)
	result := make([]string, len(As))
	for i, a := range As {
		result[i] = a.Record
	}
	if s.db.Error != nil {
		return nil, s.db.Error
	}
	return result, nil
}

func (s *sqlitedb) CreateA(domain string, records []string) error {
	As := make([]model.A, len(records))
	for i, record := range records {
		As[i] = model.A{
			Domain:     domain,
			RecordType: "A",
			Record:     record,
		}
	}
	s.db.Create(&As)
	if s.db.Error != nil {
		return s.db.Error
	}
	return nil
}

func (s *sqlitedb) UpdateA(domain string, records []string) error {
	s.db.Where("domain = ?", domain).Delete(&model.A{})
	if s.db.Error != nil {
		return s.db.Error
	}
	As := make([]model.A, len(records))
	for i, record := range records {
		As[i] = model.A{
			Domain:     domain,
			RecordType: "A",
			Record:     record,
		}
	}
	s.db.Create(&As)
	if s.db.Error != nil {
		return s.db.Error
	}
	return nil
}

func (s *sqlitedb) DeleteA(domain string) error {
	s.db.Unscoped().Where("domain = ?", domain).Delete(&model.A{})
	if s.db.Error != nil {
		return s.db.Error
	}
	return nil
}
