package types

type PersistentProcess string
type FloatingProcess string

type ProcessPlugin[_ PersistentProcess | FloatingProcess] struct {
	Handler func(*Anigo) error
	Test    func() []string
}

type ServicePlugin struct {
	Handler func(mod string, params ...string) (interface{}, bool)
	Test    func() []string
	Solvers []string
}

type ProviderPlugin struct {
	Handler func(domain string, url string) (directUrl string, ok bool)
	Test    func() []string
	Solvers []string
}

type Anigo struct {
	Global map[string]interface{}
}
