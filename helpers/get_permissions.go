package helpers

import (
	"context"
	"fmt"
	"log"
	"strings"

	iam "google.golang.org/api/iam/v1"
)

func GetPermissions(role string) ([]string, error) {
	if !strings.Contains(role, "roles/") {
		role = "roles/" + role
	}

	ctx := context.Background()

	iamService, err := iam.NewService(ctx)
	if err != nil {
		log.Panic(err)
	}

	resp, err := iamService.Roles.Get(role).Context(ctx).Do()
	if err != nil {
		return []string{}, fmt.Errorf("failed to find any permissions for the given role: %s", role)
	} else {
		return resp.IncludedPermissions, nil
	}
}
