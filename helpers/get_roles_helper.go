package helpers

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/api/iam/v1"
)

func GetRolesHelper(userPerm []string) ([]*iam.Role, error) {
	var roleList []*iam.Role
	var counter int

	ctx := context.Background()

	iamService, err := iam.NewService(ctx)
	if err != nil {
		log.Panic(err)
	}

	req := iamService.Roles.List().View("FULL")

	if err := req.Pages(ctx, func(page *iam.ListRolesResponse) error {
		for _, role := range page.Roles {
			counter = 0
			for _, googlePerm := range role.IncludedPermissions {
				for _, perm := range userPerm {
					if perm == googlePerm {
						counter++
					}
				}

			}
			if counter == len(userPerm) {
				roleList = append(roleList, role)
			}
		}
		return nil
	}); err != nil {
		log.Panic(err)
	}

	if len(roleList) > 0 {
		return roleList, nil
	} else {
		return []*iam.Role{}, fmt.Errorf("No Roles with the given Permission(s): %s", userPerm)
	}
}
