package helpers

import (
	"testing"
)

func TestPermissionTrue(t *testing.T) {

	got, _ := AllMatchedRoles([]string{"resourcemanager.folders.create"})

	if got == nil {
		t.Errorf("Role(s) should be returned for this Permission")
	}
}

func TestPermissionMultiple(t *testing.T) {

	got, _ := AllMatchedRoles([]string{"compute.networks.get", "compute.zones.list"})

	if got == nil {
		t.Errorf("Role(s) should be returned for this Permission")
	}
}

func TestPermissionFalse(t *testing.T) {

	_, err := AllMatchedRoles([]string{"nonsense.permission"})

	if err == nil {
		t.Errorf("Role(s) should not be returned for this Permission")
	}
}
