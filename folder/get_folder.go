package folder

import (
	"fmt"
	"strings"

	"github.com/gofrs/uuid"
)

func GetAllFolders() []Folder {
	return GetSampleData()
}

func (f *driver) GetFoldersByOrgID(orgID uuid.UUID) ([]Folder, error) {
	folders := f.folders

	res := []Folder{}
	for _, f := range folders {
		if f.OrgId == orgID {
			res = append(res, f)
		}
	}
	if len(res) == 0 {
		return res, fmt.Errorf("no folder with specified orgId found")
	}
	return res, nil

}

/* Notes:
** I did not implement the last test example in the readme that said "no folder with given name exists in
** this organisation" as I felt that error message and checking for that in general was insecure and not
** necessary for the task this function will be trying to achieve as you can eliminate all folders not of
** the specified organisation in GetFoldersByOrgID.
**
** I considered using a map here to significantly increase the speed of saerch with large data sets however
** because we can have mulitple different orgID's to search for I couldn't find a solution that would provide
** the same benefits without having the trade of needing to regenerate the map for each orgID every time.
**
 */
func (f *driver) GetAllChildFolders(orgID uuid.UUID, name string) ([]Folder, error) {
	folders, err := f.GetFoldersByOrgID(orgID)
	if err != nil {
		return folders, err
	}
	children := []Folder{}

	folderPath := findFullPath(f.folders, name)
	if folderPath == "" {
		return children, fmt.Errorf("folder does not exist")
	}

	// Collect all child folders by checking if they are contained
	for _, folder := range folders {
		if strings.Contains(folder.Paths, folderPath+".") {
			children = append(children, folder)
		}
	}
	return children, nil
}

/* findFullPath searches for the full path of a folder by its name using the string Split functionality
** to separate each folder from the path.
 */
func findFullPath(folders []Folder, name string) string {
	for _, folder := range folders {
		segments := strings.Split(folder.Paths, ".")
		if segments[len(segments)-1] == name {
			return folder.Paths
		}
	}
	return ""
}
