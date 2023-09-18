package helpers

import (
	"testing"
)

func TestLpPermissionTrue(t *testing.T) {

	got, _ := PolpMatchedRoles([]string{"compute.snapshots.get"})

	if got == nil {
		t.Errorf("Role(s) should be returned for this Permission")
	}
}

func TestLpPermissionMultiple(t *testing.T) {

	got, _ := PolpMatchedRoles([]string{"resourcemanager.projects.get", "compute.vpnGateways.list", "compute.urlMaps.get", "compute.targetTcpProxies.get", "compute.targetPools.list"})

	if got == nil {
		t.Errorf("Role(s) should be returned for this Permission")
	}
}

func TestLpPermissionFalse(t *testing.T) {

	_, err := PolpMatchedRoles([]string{"nonsense.permission"})

	if err == nil {
		t.Errorf("Role(s) should not be returned for this Permission")
	}
}
