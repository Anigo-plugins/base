package types

type PersistentProcess string
type FloatingProcess string

type ProcessPlugin[_ PersistentProcess | FloatingProcess] struct {
	Handler func(*Anigo) error
	Test    func() []string
}

type ServicePlugin struct {
	Handler func(string, ...string) (string, bool)
	Test    func() []string
	Solvers []string
}

type ProviderPlugin struct {
	Handler func(string, string) (string, bool)
	Test    func() []string
	Solvers []string
}

type Anigo struct {
	Global map[string]interface{}
}
