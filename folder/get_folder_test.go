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
			// Sample.json has 30 instances of the "fast-watchmen" folder copied here to test.
			name:    "bulk-test functionality",
			folder:  "fast-watchmen",
			orgID:   uuid.FromStringOrNil(folder.DefaultOrgID),
			folders: folder.GetAllFolders(),
			want: testingReturn{
				folders: []folder.Folder{
					{
						Name:  "fast-watchmen",
						OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
						Paths: "noble-vixen.fast-watchmen",
					},
					{
						Name:  "full-weapon-x",
						OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
						Paths: "noble-vixen.fast-watchmen.full-weapon-x",
					},
					{
						Name:  "honest-greymalkin",
						OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
						Paths: "noble-vixen.fast-watchmen.full-weapon-x.honest-greymalkin",
					},
					{
						Name:  "settled-copperhead",
						OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
						Paths: "noble-vixen.fast-watchmen.full-weapon-x.honest-greymalkin.settled-copperhead",
					},
					{
						Name:  "flexible-the-hunter",
						OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
						Paths: "noble-vixen.fast-watchmen.full-weapon-x.honest-greymalkin.flexible-the-hunter",
					},
					{
						Name:  "dashing-mirage",
						OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
						Paths: "noble-vixen.fast-watchmen.full-weapon-x.honest-greymalkin.dashing-mirage",
					},
					{
						Name:  "probable-oracle",
						OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
						Paths: "noble-vixen.fast-watchmen.full-weapon-x.probable-oracle",
					},
					{
						Name:  "amazing-bubbles",
						OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
						Paths: "noble-vixen.fast-watchmen.full-weapon-x.probable-oracle.amazing-bubbles",
					},
					{
						Name:  "prompt-flaberella",
						OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
						Paths: "noble-vixen.fast-watchmen.full-weapon-x.probable-oracle.prompt-flaberella",
					},
					{
						Name:  "strong-elongated",
						OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
						Paths: "noble-vixen.fast-watchmen.full-weapon-x.probable-oracle.strong-elongated",
					},
					{
						Name:  "trusty-violator",
						OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
						Paths: "noble-vixen.fast-watchmen.full-weapon-x.probable-oracle.trusty-violator",
					},
					{
						Name:  "deciding-famine",
						OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
						Paths: "noble-vixen.fast-watchmen.deciding-famine",
					},
					{
						Name:  "mint-dream",
						OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
						Paths: "noble-vixen.fast-watchmen.deciding-famine.mint-dream",
					},
					{
						Name:  "giving-stilt-man",
						OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
						Paths: "noble-vixen.fast-watchmen.deciding-famine.mint-dream.giving-stilt-man",
					},
					{
						Name:  "mutual-cyclone",
						OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
						Paths: "noble-vixen.fast-watchmen.deciding-famine.mint-dream.mutual-cyclone",
					},
					{
						Name:  "modern-silver-sable",
						OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
						Paths: "noble-vixen.fast-watchmen.deciding-famine.mint-dream.modern-silver-sable",
					},
					{
						Name:  "main-man-wolf",
						OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
						Paths: "noble-vixen.fast-watchmen.deciding-famine.mint-dream.main-man-wolf",
					},
					{
						Name:  "rational-gauntlet",
						OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
						Paths: "noble-vixen.fast-watchmen.deciding-famine.rational-gauntlet",
					},
					{
						Name:  "evolved-bastion",
						OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
						Paths: "noble-vixen.fast-watchmen.deciding-famine.rational-gauntlet.evolved-bastion",
					},
					{
						Name:  "growing-menace",
						OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
						Paths: "noble-vixen.fast-watchmen.growing-menace",
					},
					{
						Name:  "super-cobweb",
						OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
						Paths: "noble-vixen.fast-watchmen.growing-menace.super-cobweb",
					},
					{
						Name:  "perfect-vanisher",
						OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
						Paths: "noble-vixen.fast-watchmen.growing-menace.super-cobweb.perfect-vanisher",
					},
					{
						Name:  "settling-hobgoblin",
						OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
						Paths: "noble-vixen.fast-watchmen.settling-hobgoblin",
					},
					{
						Name:  "super-stunner",
						OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
						Paths: "noble-vixen.fast-watchmen.settling-hobgoblin.super-stunner",
					},
					{
						Name:  "fine-haven",
						OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
						Paths: "noble-vixen.fast-watchmen.settling-hobgoblin.super-stunner.fine-haven",
					},
					{
						Name:  "huge-witchblade",
						OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
						Paths: "noble-vixen.fast-watchmen.settling-hobgoblin.huge-witchblade",
					},
					{
						Name:  "fine-shredder",
						OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
						Paths: "noble-vixen.fast-watchmen.settling-hobgoblin.huge-witchblade.fine-shredder",
					},
					{
						Name:  "noted-lady-bullseye",
						OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
						Paths: "noble-vixen.fast-watchmen.settling-hobgoblin.noted-lady-bullseye",
					},
					{
						Name:  "welcomed-crazy",
						OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
						Paths: "noble-vixen.fast-watchmen.settling-hobgoblin.noted-lady-bullseye.welcomed-crazy",
					},
					{
						Name:  "nearby-beetle",
						OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
						Paths: "noble-vixen.fast-watchmen.settling-hobgoblin.noted-lady-bullseye.nearby-beetle",
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
