package types

import (
	"testing"
)

func Test_A2_GetServicesNames(t *testing.T) {
	if _, ok := Anigo.GetServiceNames(); !ok {
		t.Errorf("Expected %t, got %t", false, true)

	}

	if _, ok := Anigo2.GetServiceNames(); ok {
		t.Errorf("Expected %t, got %t", true, false)

	}

}
