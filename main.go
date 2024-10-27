package main

import (
	"fmt"
	"strings"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
)

func main() {
	// Change orgID here to work with different sets of sample data.
	orgID := uuid.FromStringOrNil(folder.DefaultOrgID)

	res := folder.GetAllFolders()

	folderDriver := folder.NewDriver(res)
	// example usage
	// orgFolder, error := folderDriver.GetFoldersByOrgID(orgID)

	// if error == nil {
	// 	folder.PrettyPrint(res)
	// 	fmt.Printf("\n Folders for orgID: %s", orgID)
	// 	folder.PrettyPrint(orgFolder)
	// }

	// fmt.Println(error)

	// input handling for component 1 & 2
	// adapted from my work here: https://github.com/bucL/sc-interns-takehometask/blob/main/main.go

	var input string
	for input != "exit" {
		fmt.Println("Type 1 for component 1 and 2 for component 2 or exit to close")
		fmt.Scan(&input)
		if input != "exit" {
			switch input {
			case "1":
				var folderName string
				fmt.Println("Enter folder name")
				fmt.Scan(&folderName)
				folderName = strings.TrimSpace(folderName)
				folders, err := folderDriver.GetAllChildFolders(orgID, folderName)
				if err != nil {
					fmt.Printf("%s\n", err)
				} else {
					folder.PrettyPrint(folders)
					fmt.Println("")
				}
			case "2":
				var source string
				var destination string
				fmt.Println("Enter source folder")
				fmt.Scan(&source)
				source = strings.TrimSpace(source)
				fmt.Println("Enter destination folder")
				fmt.Scan(&destination)
				destination = strings.TrimSpace(destination)

				folders, err := folderDriver.MoveFolder(orgID, source, destination)
				if err != nil {
					fmt.Printf("%s\n", err)
				} else {
					folder.PrettyPrint(folders)
					fmt.Println("")
				}
			default:
				fmt.Println("Invalid Option Try Again")
			}

		}
	}
}
