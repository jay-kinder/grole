package helpers

import (
	"context"
	"fmt"

	iam "google.golang.org/api/iam/v1"
)

func GetRolesHelper(userPerm []string) ([]*iam.Role, error) {
	var roleList []*iam.Role

	// Build a map of user permissions for O(1) lookup
	userPermMap := make(map[string]bool, len(userPerm))
	for _, perm := range userPerm {
		userPermMap[perm] = true
	}

	ctx := context.Background()

	iamService, err := iam.NewService(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create IAM service: %w", err)
	}

	req := iamService.Roles.List().View("FULL")

	if err := req.Pages(ctx, func(page *iam.ListRolesResponse) error {
		for _, role := range page.Roles {
			counter := 0
			for _, googlePerm := range role.IncludedPermissions {
				if userPermMap[googlePerm] {
					counter++
				}
			}
			if counter == len(userPerm) {
				roleList = append(roleList, role)
			}
		}
		return nil
	}); err != nil {
		return nil, fmt.Errorf("failed to list roles: %w", err)
	}

	if len(roleList) > 0 {
		return roleList, nil
	}
	return []*iam.Role{}, fmt.Errorf("no roles with the given permission(s): %s", userPerm)
}
