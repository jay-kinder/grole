package helpers

func PolpMatchedRoles(userPerm []string) ([]string, error) {
	matchedPerms, err := GetRolesHelper(userPerm)
	if err != nil {
		return []string{}, err
	}

	if len(matchedPerms) == 0 {
		return []string{}, nil
	}

	// Initialize with the first role's permission count
	amountOfPerms := len(matchedPerms[0].IncludedPermissions)
	leastPrivRoleList := []string{matchedPerms[0].Name}

	for _, role := range matchedPerms[1:] {
		rolePermCount := len(role.IncludedPermissions)
		if rolePermCount < amountOfPerms {
			amountOfPerms = rolePermCount
			leastPrivRoleList = []string{role.Name}
		} else if rolePermCount == amountOfPerms {
			leastPrivRoleList = append(leastPrivRoleList, role.Name)
		}
	}
	return leastPrivRoleList, nil
}
