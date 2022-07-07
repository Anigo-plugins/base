package types

import (
	"testing"
)

var _DomainTests = map[string]string{
	"www.youtube.com": "https://www.youtube.com/watch?v=Wc6GGsf5oxE",
	"github.com":      "https://github.com/FlamesX-128",
}

func Test_A3_GetDomain(t *testing.T) {
	for k, v := range _DomainTests {
		if d := Anigo.GetDomain(v); d != k {
			t.Errorf("Expected %s, got %s", v, d)

		}

	}

}
