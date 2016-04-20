package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/codegangsta/cli"
)

var rolesFolders = [7]string{"tasks", "templates", "vars", "files", "meta", "handlers", "defaults"}

func main() {
	app := cli.NewApp()
	app.Name = "utilsible"
	app.Usage = "Command line utilities for Ansible"
	app.Commands = []cli.Command{
		{
			Name:    "roles",
			Aliases: []string{"r"},
			Usage:   "Actions on roles",
			Subcommands: []cli.Command{
				{
					Name:   "add",
					Usage:  "add a new role",
					Action: addNewRole,
				},
				{
					Name:   "clean",
					Usage:  "clean an existing role",
					Action: cleanRole,
				},
				{
					Name:   "lint",
					Usage:  "perform linting over a role",
					Action: lintRole,
				},
			},
		},
	}

	app.Run(os.Args)
}

func cleanRole(c *cli.Context) {
	println("Cleaning role")
}

func lintRole(c *cli.Context) {
	println("Linting role")
}

func addNewRole(c *cli.Context) {
	if len(c.Args()) == 0 {
		println("No name of new role passed")
		return
	}

	roleName := c.Args()[0]
	fmt.Printf("Adding new role %s\n", roleName)

	//Permissions for the new files and folders
	var dirMode os.FileMode = 0775

	//Check if we are in root Ansible folder or already in roles
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	if !strings.Contains(wd, "roles") {
		//Check if the roles folder already exists or create it
		_ = os.Mkdir("roles", dirMode)
		err = os.Chdir("roles")
		if err != nil {
			log.Fatal(err)
		}
	}

	//Create role folder
	err = os.Mkdir(roleName, dirMode)
	if err != nil {
		log.Fatal(err)
	}

	//Change dir to the newly created
	err = os.Chdir(roleName)
	if err != nil {
		log.Fatal(err)
	}

	//Create all roles subfolders
	for _, role := range rolesFolders {
		err := os.Mkdir(role, dirMode)
		if err != nil {
			log.Fatal(err)
		}
	}
}
