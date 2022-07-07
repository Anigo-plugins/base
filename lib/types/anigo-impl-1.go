package types

func (a *Anigo) GetProviders() ([]ProviderPlugin, bool) {
	if prov, ok := a.Global["Providers"]; ok {
		if p, ok := prov.([]ProviderPlugin); ok {

			return p, true
		}

	}

	return []ProviderPlugin{}, false
}

func (a *Anigo) GetProviderNames() ([]string, bool) {
	if prov, ok := a.GetProviders(); ok {
		var solvers []string

		for _, p := range prov {
			solvers = append(solvers, p.Solvers...)

		}

		return solvers, true
	}

	return []string{}, false
}
