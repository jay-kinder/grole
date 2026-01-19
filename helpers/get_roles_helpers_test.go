package helpers

import (
	"testing"
)

func TestRhPermTrue(t *testing.T) {

	got, err := GetRolesHelper([]string{"compute.networks.get"})

	if err != nil {
		t.Errorf("Error should not be returned for valid permission: %v", err)
	}

	if len(got) == 0 {
		t.Errorf("Permissions should be returned for this Role")
	}
}

func TestRhPermFalse(t *testing.T) {

	_, err := GetRolesHelper([]string{"nonsense.permission"})

	if err == nil {
		t.Errorf("Permissions should not be returned for this Role")
	}
}
