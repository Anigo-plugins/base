package types

func (a *Anigo) GetServices() ([]ServicePlugin, bool) {
	if svc, ok := a.Global["Services"]; ok {
		if s, ok := svc.([]ServicePlugin); ok {

			return s, true
		}

	}

	return []ServicePlugin{}, false
}

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
