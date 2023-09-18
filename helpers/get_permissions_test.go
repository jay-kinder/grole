package helpers

import (
	"testing"
)

func TestRoleTruePrefix(t *testing.T) {

	got, _ := GetPermissions("roles/compute.admin")

	if got == nil {
		t.Errorf("Permissions should be returned for this Role")
	}
}

func TestRoleTrueNoPrefix(t *testing.T) {

	got, _ := GetPermissions("riscconfigs.admin")

	if got == nil {
		t.Errorf("Permissions should be returned for this Role")
	}
}

func TestRoleFalse(t *testing.T) {

	_, err := GetPermissions("nonsense.role")

	if err == nil {
		t.Errorf("Permissions should not be returned for this Role")
	}
}
