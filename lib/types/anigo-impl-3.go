package types

import (
	"log"
	"strings"

	"github.com/Anigo-plugins/base/lib/tools"
)

func (a *Anigo) GetDomain(url string) (string, bool) {
	if !strings.HasPrefix(url, "https://") {

		return url, false
	}

	u := strings.Replace(url, "https://", "", -1)

	return u[:strings.Index(u, "/")], true
}

// // --------------------------------------------------- \\ \\

func (a *Anigo) GetDirectUrl(url string) (string, bool) {
	providers, _ := a.GetProviders()

	if len(providers) == 0 {
		log.Println("Providers not found.")

	}

	if domain, ok := a.GetDomain(url); ok {
		for _, prov := range providers {
			if !tools.Contains(prov.Solvers, domain) {

				continue
			}

			if url, ok := prov.Handler(domain, url); ok {

				return url, true
			}

		}

	}

	return url, false
}
