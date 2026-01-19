package helpers

func AllMatchedRoles(userPerm []string) ([]string, error) {
	matchedPerms, err := GetRolesHelper(userPerm)
	if err != nil {
		return []string{}, err
	}

	rolesList := make([]string, 0, len(matchedPerms))
	for _, role := range matchedPerms {
		rolesList = append(rolesList, role.Name)
	}
	return rolesList, nil
}
