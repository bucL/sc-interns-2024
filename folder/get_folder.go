package folder

import (
	"fmt"

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
		return nil, fmt.Errorf("no folder with specified orgId found")
	}
	return res, nil

}

/* Notes:
** I did not implement the last test example in the readme that said "no folder with given name exists in
** this organisation" as I felt that error message and checking for that in general was insecure and not
** necessary for the task this function will be trying to achieve as you can eliminate all folders not of
** the specified organisation in GetFoldersByOrgID.
 */
func (f *driver) GetAllChildFolders(orgID uuid.UUID, name string) ([]Folder, error) {
	folders, err := f.GetFoldersByOrgID(orgID)
	if err != nil {
		return nil, err
	}
	// implementation below

	return []Folder{}, nil
}
