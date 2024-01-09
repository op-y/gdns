package cache

type Cache interface {
	Name() string
	QueryA(string) ([]string, error)
	UpdateA(string, []string) error
	DeleteA(string) error
}
