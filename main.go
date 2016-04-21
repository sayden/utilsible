package main

import (
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
					Action: getAddNewRoleFunc,
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

func getAddNewRoleFunc(c *cli.Context) {
	addNewRole(c.Args())
}

// isRoot checks if we are in root Ansible folder or already in roles
func isRoot() bool {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	return !strings.Contains(wd, "roles")
}
