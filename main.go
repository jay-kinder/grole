package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/akamensky/argparse"
	"github.com/jay-kinder/grole/helpers"
)

var version = "v1.0.0"

func main() {
	parser := argparse.NewParser("grole", "☁️  grole is a command line tool for establishing which Google Cloud Roles contain a given permission, and also which permissions a given Role has.")

	roleFlag := parser.String("r", "role", &argparse.Options{Required: false, Help: "Provide a role name to see all the permissions it has"})

	permFlag := parser.StringList("p", "perm", &argparse.Options{Required: false, Help: "Provide permission(s) and see which role(s) contain both this permission(s) and the smallest number of other permissions (helps to follow the principle of least privilege)"})

	allFlag := parser.Flag("", "all", &argparse.Options{Required: false, Help: "Provide permission(s) with [-p | --perm ] and see all role(s) that contain this permission(s)"})

	versionFlag := parser.Flag("v", "version", &argparse.Options{Required: false, Help: "Get currently installed version of grole"})

	err := parser.Parse(os.Args)

	if err != nil || len(os.Args) == 1 {
		fmt.Print(parser.Usage(err))
	}

	switch {
	case *roleFlag != "":
		permissionsList, err := helpers.GetPermissions(*roleFlag)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(strings.Join(permissionsList, "\n"))

	case len(*permFlag) != 0:
		if *allFlag {
			rolesList, err := helpers.AllMatchedRoles(*permFlag)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(strings.Join(rolesList, "\n"))
		} else {
			leastPrivRolesList, err := helpers.PolpMatchedRoles(*permFlag)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(strings.Join(leastPrivRolesList, "\n"))
		}

	case *versionFlag:
		fmt.Println(version)
	}
}
