package types

import (
	"log"

	"github.com/Anigo-plugins/base/lib/tools"
)

func (a *Anigo) GetServices() ([]ServicePlugin, bool) {
	if svc, ok := a.Global["Services"]; ok {
		if s, ok := svc.([]ServicePlugin); ok {

			return s, true
		}

	}

	return []ServicePlugin{}, false
}

// // --------------------------------------------------- \\ \\

func (a *Anigo) GetServiceNames() ([]string, bool) {
	if svc, ok := a.GetServices(); ok {
		var solvers []string

		for _, s := range svc {
			solvers = append(solvers, s.Solvers...)

		}

		return solvers, true
	}

	return []string{}, false
}

// // --------------------------------------------------- \\ \\

func (a *Anigo) UseService(svc string, mod string, params ...string) (interface{}, bool) {
	services, _ := a.GetServices()

	if len(services) == 0 {
		log.Println("Services not found.")

	}

	for _, s := range services {
		if !tools.Contains(s.Solvers, svc) {

			continue
		}

		if res, ok := s.Handler(mod, params...); ok {

			return res, true
		}

	}

	return []interface{}{}, false
}
