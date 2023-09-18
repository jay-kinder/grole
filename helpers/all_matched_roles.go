package helpers

func AllMatchedRoles(userPerm []string) ([]string, error) {
	matchedPerms, err := GetRolesHelper(userPerm)
	rolesList := []string{}

	if err != nil {
		return []string{}, err
	} else {
		for _, role := range matchedPerms {
			rolesList = append(rolesList, role.Name)
		}
		return rolesList, nil
	}
}
