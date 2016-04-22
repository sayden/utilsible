package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

func createTextFile(role string, t string, subfolder string){
	f, err := os.Create(t)
	if err != nil {
		log.Fatal(err)
	}

	tmpl, err := template.ParseFiles(os.Getenv("UTILSIBLE") + "/templates/roles/" + t)
	if err != nil {
		log.Fatal(err)
	}

	_role := Role{
		Role: role,
		Subfolder:subfolder,
	}

	err = tmpl.Execute(f, _role)
	if err != nil {
		log.Fatal(err)
	}
}

// Role is a struct to hold the information about a role that is being
// created and the current subfolder
type Role struct {
	Role      string
	Subfolder string
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

	createTextFile(roleName, "README.md", "")

	//Create all roles subfolders
	for _, subfolder := range rolesFolders {
		err := os.Mkdir(subfolder, dirMode)
		if err != nil {
			log.Fatal(err)
		}

		//Change to the new folder
		cur, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}

		err = os.Chdir(subfolder)
		if err != nil {
			log.Fatal(err)
		}

		createTextFile(roleName, "main.yml", subfolder)

		//Come to the previous folder
		err = os.Chdir(cur)
		if err != nil {
			log.Fatal(err)
		}
	}
}
