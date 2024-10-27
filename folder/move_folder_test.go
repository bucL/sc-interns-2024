package folder_test

import (
	"fmt"
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_folder_MoveFolder(t *testing.T) {
	t.Parallel()
	tests := [...]struct {
		name        string
		orgID       uuid.UUID
		source      string
		destination string
		folders     []folder.Folder
		want        testingReturn
	}{
		{
			name:        "success path",
			orgID:       uuid.FromStringOrNil(folder.DefaultOrgID),
			source:      "bravo",
			destination: "charlie",
			folders: []folder.Folder{
				{
					Name:  "Alpha",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
					Paths: "alpha",
				}, {
					Name:  "charlie",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
					Paths: "charlie",
				},
				{
					Name:  "bravo",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
					Paths: "alpha.bravo",
				},
				{
					Name:  "delta",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
					Paths: "alpha.bravo.delta",
				},
			},
			want: testingReturn{
				folders: []folder.Folder{
					{
						Name:  "Alpha",
						OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
						Paths: "alpha",
					}, {
						Name:  "charlie",
						OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
						Paths: "charlie",
					},
					{
						Name:  "bravo",
						OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
						Paths: "alpha.charlie.bravo",
					},
					{
						Name:  "delta",
						OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
						Paths: "alpha.charlie.bravo.delta",
					},
				},
				err: nil,
			},
		},
		{
			name:        "invalid orgId",
			orgID:       uuid.FromStringOrNil(folder.AlternativeOrgID),
			source:      "bravo",
			destination: "charlie",
			folders: []folder.Folder{
				{
					Name:  "Alpha",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
					Paths: "alpha",
				}, {
					Name:  "charlie",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
					Paths: "charlie",
				},
				{
					Name:  "bravo",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
					Paths: "alpha.bravo",
				},
				{
					Name:  "delta",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
					Paths: "alpha.bravo.delta",
				},
			},
			want: testingReturn{
				folders: []folder.Folder{},
				err:     fmt.Errorf("folder does not exist"),
			},
		},
		{
			name:        "move folder to itself",
			orgID:       uuid.FromStringOrNil(folder.DefaultOrgID),
			source:      "bravo",
			destination: "bravo",
			folders: []folder.Folder{
				{
					Name:  "Alpha",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
					Paths: "alpha",
				}, {
					Name:  "charlie",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
					Paths: "charlie",
				},
				{
					Name:  "bravo",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
					Paths: "alpha.bravo",
				},
				{
					Name:  "delta",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
					Paths: "alpha.bravo.delta",
				},
			},
			want: testingReturn{
				folders: []folder.Folder{},
				err:     fmt.Errorf("cannot move folder to itself"),
			},
		},
		{
			name:        "move folder to child of itself",
			orgID:       uuid.FromStringOrNil(folder.DefaultOrgID),
			source:      "bravo",
			destination: "delta",
			folders: []folder.Folder{
				{
					Name:  "Alpha",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
					Paths: "alpha",
				}, {
					Name:  "charlie",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
					Paths: "charlie",
				},
				{
					Name:  "bravo",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
					Paths: "alpha.bravo",
				},
				{
					Name:  "delta",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
					Paths: "alpha.bravo.delta",
				},
			},
			want: testingReturn{
				folders: []folder.Folder{},
				err:     fmt.Errorf("cannot move folder to child of itself"),
			},
		},
		{
			name:        "move folder to itself",
			orgID:       uuid.FromStringOrNil(folder.DefaultOrgID),
			source:      "bravo",
			destination: "bravo",
			folders: []folder.Folder{
				{
					Name:  "Alpha",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
					Paths: "alpha",
				}, {
					Name:  "charlie",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
					Paths: "charlie",
				},
				{
					Name:  "bravo",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
					Paths: "alpha.bravo",
				},
				{
					Name:  "delta",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
					Paths: "alpha.bravo.delta",
				},
			},
			want: testingReturn{
				folders: []folder.Folder{},
				err:     fmt.Errorf("cannot move folder to itself"),
			},
		},
		{
			name:        "source doesn't exist",
			orgID:       uuid.FromStringOrNil(folder.DefaultOrgID),
			source:      "foxtrot",
			destination: "bravo",
			folders: []folder.Folder{
				{
					Name:  "Alpha",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
					Paths: "alpha",
				}, {
					Name:  "charlie",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
					Paths: "charlie",
				},
				{
					Name:  "bravo",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
					Paths: "alpha.bravo",
				},
				{
					Name:  "delta",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
					Paths: "alpha.bravo.delta",
				},
			},
			want: testingReturn{
				folders: []folder.Folder{},
				err:     fmt.Errorf("source folder doesn't exist"),
			},
		},
		{
			name:        "destination doesn't exist",
			orgID:       uuid.FromStringOrNil(folder.DefaultOrgID),
			source:      "bravo",
			destination: "bravo",
			folders: []folder.Folder{
				{
					Name:  "Alpha",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
					Paths: "alpha",
				}, {
					Name:  "charlie",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
					Paths: "charlie",
				},
				{
					Name:  "bravo",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
					Paths: "alpha.bravo",
				},
				{
					Name:  "delta",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
					Paths: "alpha.bravo.delta",
				},
			},
			want: testingReturn{
				folders: []folder.Folder{},
				err:     fmt.Errorf("detination folder doesn't exist"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			get, err := f.MoveFolder(tt.orgID, tt.source, tt.destination)
			assert := assert.New(t)

			assert.Equal(tt.want.folders, get)
			assert.Equal(tt.want.err, err)
		})
	}
}
