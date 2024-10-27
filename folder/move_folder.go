package folder

import (
	"fmt"
	"strings"

	"github.com/gofrs/uuid"
)

func (f *driver) MoveFolder(orgID uuid.UUID, source string, dst string) ([]Folder, error) {
	// Retrieve all folders for the specified organization
	folders := f.folders

	sourcePath := findFullPath(folders, source)
	destinationPath := findFullPath(folders, dst)

	// Check OrgID of destination because checking orgID of source is done when calling GetAllChildFolders below.
	for folder := range folders {
		if folders[folder].Paths == destinationPath {
			if folders[folder].OrgId != orgID {
				return []Folder{}, fmt.Errorf("no folder with specified orgId found")
			} else {
				break
			}
		}
	}

	if sourcePath == "" {
		return []Folder{}, fmt.Errorf("source folder doesn't exist")
	}
	if destinationPath == "" {
		return []Folder{}, fmt.Errorf("destination folder doesn't exist")
	}
	if sourcePath == destinationPath {
		return []Folder{}, fmt.Errorf("cannot move folder to itself")
	}
	if strings.HasPrefix(destinationPath, sourcePath+".") {
		return []Folder{}, fmt.Errorf("cannot move folder to child of itself")
	}

	newSourcePath := destinationPath + "." + source

	// Use GetAllChildFolders to get children of the source folder
	children, err := f.GetAllChildFolders(orgID, source)
	if err != nil {
		return []Folder{}, err
	}

	// Update paths for the source folder and each child
	for i := range folders {
		// Update the source folder path
		if folders[i].Paths == sourcePath {
			folders[i].Paths = newSourcePath
		} else {
			// If the folder is a child of the source folder, update its path
			for _, child := range children {
				if folders[i].Paths == child.Paths {
					folders[i].Paths = strings.Replace(folders[i].Paths, sourcePath, newSourcePath, 1)
					break
				}
			}
		}
	}

	return folders, nil // Return the updated folders
}
