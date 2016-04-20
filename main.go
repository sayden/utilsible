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

//addNewRole adds a new folder with the role passed as argument. It creates
//roles folder as neccesary if it doesn't exists yet
func addNewRole(args []string) {
	if len(args) == 0 {
		println("No name of new role passed")
		return
	}

	roleName := args[0]
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

	createReadmeFile(roleName)

	//Create all roles subfolders
	for _, role := range rolesFolders {
		err := os.Mkdir(role, dirMode)
		if err != nil {
			log.Fatal(err)
		}

		//Change to the new folder
		cur, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}

		err = os.Chdir(role)
		if err != nil {
			log.Fatal(err)
		}
		createMainFileWithComments(roleName, role)

		//Come to the previous folder
		err = os.Chdir(cur)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func createMainFileWithComments(r string, subfolder string) {
	f, err := os.Create("main.yml")
	if err != nil {
		log.Fatal(err)
	}

	fileContent := fmt.Sprintf("# roles/%s/%s/main.yml\n", r, subfolder)

	_, err = f.WriteString(fileContent)
	if err != nil {
		log.Fatal(err)
	}
}

func createReadmeFile(r string) {
	f, err := os.Create("README.md")
	if err != nil {
		log.Fatal(err)
	}

	_, err = f.WriteString(fmt.Sprintf("# Role: %s\n", r))
	if err != nil {
		log.Fatal(err)
	}
}
