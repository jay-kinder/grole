package helpers

import (
	"testing"
)

func TestRhPermTrue(t *testing.T) {

	got, _ := GetRolesHelper([]string{"compute.networks.get"})

	if got == nil {
		t.Errorf("Permissions should be returned for this Role")
	}
}

func TestRhPermFalse(t *testing.T) {

	_, err := GetRolesHelper([]string{"nonsense.permission"})

	if err == nil {
		t.Errorf("Permissions should not be returned for this Role")
	}
}
