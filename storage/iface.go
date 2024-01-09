package storage

type Storage interface {
	Name() string
	QueryA(string) ([]string, error)
	CreateA(string, []string) error
	UpdateA(string, []string) error
	DeleteA(string) error
}
