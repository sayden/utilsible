package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

//createReadmeFile creates a Readme file for the root folder of every role
func createReadmeFile(r string) {
	f, err := os.Create("README.md")
	if err != nil {
		log.Fatal(err)
	}

	role := Role{
		Role: r,
	}

	tmpl, err := template.ParseFiles("../../templates/README.md")
	if err != nil {
		log.Fatal(err)
	}

	err = tmpl.Execute(f, role)
}

// Role is a struct to hold the information about a role that is being
// created and the current subfolder
type Role struct {
	Role      string
	Subfolder string
}

// createMainFileWithComments creates a main.yml file within specified subfolder
func createMainFileWithComments(r string, subfolder string) {
	f, err := os.Create("main.yml")
	if err != nil {
		log.Fatal(err)
	}

	tmpl, err := template.ParseFiles("../../../templates/rolesMain.yml")
	if err != nil {
		log.Fatal(err)
	}

	tmplRole := Role{
		Role:      r,
		Subfolder: subfolder,
	}
	err = tmpl.Execute(f, tmplRole)
	if err != nil {
		log.Fatal(err)
	}
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

	if isRoot() {
		//Check if the roles folder already exists or create it
		_ = os.Mkdir("roles", dirMode)
		err := os.Chdir("roles")
		if err != nil {
			log.Fatal(err)
		}
	}

	//Create role folder
	err := os.Mkdir(roleName, dirMode)
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
