package folder

import (
	"fmt"
	"strings"

	"github.com/gofrs/uuid"
)

func (f *driver) MoveFolder(orgID uuid.UUID, source string, dst string) ([]Folder, error) {

	updatedFolders := []Folder{}

	folders, err := f.GetFoldersByOrgID(orgID)
	if err != nil {
		return updatedFolders, err
	}

	sourcePath := findFullPath(folders, source)
	destinationPath := findFullPath(f.folders, dst)

	if sourcePath == "" {
		return updatedFolders, fmt.Errorf("source folder doesn't exist")
	}
	if destinationPath == "" {
		return updatedFolders, fmt.Errorf("destination folder doesn't exist")
	}
	if sourcePath == destinationPath {
		return updatedFolders, fmt.Errorf("cannot move folder to itself")
	}

	if strings.HasPrefix(destinationPath, sourcePath+".") {
		return updatedFolders, fmt.Errorf("cannot move folder to child of itself")
	}
	newSourcePath := destinationPath + "." + source

	// Use GetAllChildFolders to get children of the source folder
	children, err := f.GetAllChildFolders(orgID, source)
	if err != nil {
		return updatedFolders, err
	}

	// Update paths for the source folder and each child
	for _, folder := range folders {
		if folder.Paths == sourcePath {
			// Update source folder path to newSourcePath
			folder.Paths = newSourcePath
		} else {
			// If the folder is a child of the source folder, update its path
			for _, child := range children {
				if folder.Paths == child.Paths {
					folder.Paths = strings.Replace(folder.Paths, sourcePath, newSourcePath, 1)
				}
			}
		}
		// Add the updated (or unchanged) folder to updatedFolders
		updatedFolders = append(updatedFolders, folder)
	}

	return updatedFolders, nil
}
