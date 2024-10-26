package main

import (
	"fmt"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
)

func main() {
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

	// input handling
	var input string
	for input != "exit" {
		fmt.Println("Enter a folder path")
		fmt.Scan(&input)
		if input != "exit" {
			folders, err := folderDriver.GetAllChildFolders(orgID, input)
			if err != nil {
				fmt.Println(err)
			}
			folder.PrettyPrint(folders)
		}
	}
}
