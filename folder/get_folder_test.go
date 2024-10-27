package folder_test

import (
	"fmt"
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

type testingReturn struct {
	folders []folder.Folder
	err     error
}

// feel free to change how the unit test is structured
func Test_folder_GetFoldersByOrgID(t *testing.T) {
	t.Parallel()
	tests := [...]struct {
		name    string
		orgID   uuid.UUID
		folders []folder.Folder
		want    testingReturn
	}{
		{
			name:  "Success path",
			orgID: uuid.FromStringOrNil(folder.DefaultOrgID),
			folders: []folder.Folder{
				{
					Name:  "Alpha",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
					Paths: "alpha",
				},
				{
					Name:  "bravo",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
					Paths: "alpha.bravo",
				},
				{
					Name:  "alpha",
					OrgId: uuid.FromStringOrNil(folder.AlternativeOrgID),
					Paths: "alpha",
				},
			},
			want: testingReturn{
				folders: []folder.Folder{
					{
						Name:  "Alpha",
						OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
						Paths: "alpha",
					},
					{
						Name:  "bravo",
						OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
						Paths: "alpha.bravo",
					},
				},
				err: nil,
			},
		}, {
			name:  "no folders exist in current org",
			orgID: uuid.FromStringOrNil(folder.DefaultOrgID),
			folders: []folder.Folder{
				{
					Name:  "alpha",
					OrgId: uuid.FromStringOrNil(folder.AlternativeOrgID),
					Paths: "alpha",
				},
				{
					Name:  "bravo",
					OrgId: uuid.FromStringOrNil(folder.AlternativeOrgID),
					Paths: "alpha.bravo",
				},
			},
			want: testingReturn{
				folders: []folder.Folder{},
				err:     fmt.Errorf("no folder with specified orgId found"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			get, err := f.GetFoldersByOrgID(tt.orgID)
			assert := assert.New(t)

			assert.Equal(tt.want.folders, get)
			assert.Equal(tt.want.err, err)

		})
	}
}

func Test_folder_GetChildFolders(t *testing.T) {
	t.Parallel()
	tests := [...]struct {
		name    string
		folder  string
		orgID   uuid.UUID
		folders []folder.Folder
		want    testingReturn
	}{
		{
			name:   "Success path",
			folder: "alpha",
			orgID:  uuid.FromStringOrNil(folder.DefaultOrgID),
			folders: []folder.Folder{
				{
					Name:  "Alpha",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
					Paths: "alpha",
				},
				{
					Name:  "bravo",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
					Paths: "alpha.bravo",
				},
				{
					Name:  "alpha",
					OrgId: uuid.FromStringOrNil(folder.AlternativeOrgID),
					Paths: "alpha",
				},
			},
			want: testingReturn{
				folders: []folder.Folder{
					{
						Name:  "bravo",
						OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
						Paths: "alpha.bravo",
					},
				},
				err: nil,
			},
		}, {
			// Sample.json has 4 child folders of the "fast-watchmen" folder copied here to test.
			name:    "bulk-test functionality",
			folder:  "fast-watchmen",
			orgID:   uuid.FromStringOrNil(folder.DefaultOrgID),
			folders: folder.GetAllFolders(),
			want: testingReturn{
				folders: []folder.Folder{
					{
						Name:  "full-weapon-x",
						OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
						Paths: "noble-vixen.fast-watchmen.full-weapon-x",
					},
					{
						Name:  "deciding-famine",
						OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
						Paths: "noble-vixen.fast-watchmen.deciding-famine",
					},
					{
						Name:  "growing-menace",
						OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
						Paths: "noble-vixen.fast-watchmen.growing-menace",
					},
					{
						Name:  "settling-hobgoblin",
						OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
						Paths: "noble-vixen.fast-watchmen.settling-hobgoblin",
					},
				},
				err: nil,
			},
		},
		{
			name:   "Folder exists but no child folder",
			folder: "charlie",
			orgID:  uuid.FromStringOrNil(folder.DefaultOrgID),
			folders: []folder.Folder{
				{
					Name:  "Alpha",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
					Paths: "alpha",
				}, {
					Name:  "charlie",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
					Paths: "alpha.charlie",
				},
			},
			want: testingReturn{
				folders: []folder.Folder{},
				err:     nil,
			},
		},
		{
			name:   "Folder does not exist in the organisation",
			folder: "bravo",
			orgID:  uuid.FromStringOrNil(folder.DefaultOrgID),
			folders: []folder.Folder{
				{
					Name:  "Alpha",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
					Paths: "alpha",
				}, {
					Name:  "charlie",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
					Paths: "alpha.charlie",
				},
			},
			want: testingReturn{
				folders: []folder.Folder{},
				err:     fmt.Errorf("folder does not exist"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			get, err := f.GetAllChildFolders(tt.orgID, tt.folder)
			assert := assert.New(t)

			assert.Equal(tt.want.folders, get)
			assert.Equal(tt.want.err, err)

		})
	}
}
