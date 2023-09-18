package helpers

func PolpMatchedRoles(userPerm []string) ([]string, error) {
	matchedPerms, err := GetRolesHelper(userPerm)
	amountOfPerms := 999
	leastPrivRoleList := []string{}

	if err != nil {
		return []string{}, err
	} else {
		for _, role := range matchedPerms {
			if len(role.IncludedPermissions) < amountOfPerms {
				amountOfPerms = len(role.IncludedPermissions)
				leastPrivRoleList = []string{role.Name}
			} else if len(role.IncludedPermissions) == amountOfPerms {
				leastPrivRoleList = append(leastPrivRoleList, role.Name)
			}
		}
		return leastPrivRoleList, nil
	}
}
