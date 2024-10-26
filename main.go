package main

import (
	"fmt"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
)

func main() {
	orgID := uuid.FromStringOrNil(folder.DefaultOrgID)

	res := folder.GetAllFolders()

	// example usage
	folderDriver := folder.NewDriver(res)
	orgFolder, error := folderDriver.GetFoldersByOrgID(orgID)

	if error == nil {
		folder.PrettyPrint(res)
		fmt.Printf("\n Folders for orgID: %s", orgID)
		folder.PrettyPrint(orgFolder)
	}

	fmt.Println(error)

}
