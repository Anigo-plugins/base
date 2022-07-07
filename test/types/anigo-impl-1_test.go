package types

import (
	"testing"

	"github.com/Anigo-plugins/base/lib/types"
)

var Anigo = types.Anigo{
	Global: map[string]interface{}{
		"Providers": []types.ProviderPlugin{},
		"Services":  []types.ServicePlugin{},
	},
}

var Anigo2 = types.Anigo{}

func Test_A1_GetProviderNames(t *testing.T) {
	if _, ok := Anigo.GetProviderNames(); !ok {
		t.Errorf("Expected %t, got %t", false, true)

	}

	if _, ok := Anigo2.GetProviderNames(); ok {
		t.Errorf("Expected %t, got %t", true, false)

	}

}
