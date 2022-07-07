package types

import "strings"

func (a *Anigo) GetDomain(url string) string {
	u := strings.Replace(url, "https://", "", -1)

	return u[:strings.Index(u, "/")]
}
