package folder_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	// "github.com/stretchr/testify/assert"
)

// feel free to change how the unit test is structured
func Test_folder_GetFoldersByOrgID(t *testing.T) {
	t.Parallel()
	org1ID := uuid.FromStringOrNil("00000000-0000-0000-0000-000000000001")
	org2ID := uuid.FromStringOrNil("00000000-0000-0000-0000-000000000002")
	folderData := []folder.Folder{
		{Name: "alpha", OrgId: org1ID, Paths: "alpha"},
		{Name: "bravo", Paths: "alpha.bravo", OrgId: org1ID},
		{Name: "charlie", Paths: "alpha.bravo.charlie", OrgId: org1ID},
		{Name: "delta", Paths: "alpha.delta", OrgId: org1ID},
		{Name: "echo", Paths: "echo", OrgId: org1ID},
		{Name: "foxtrot", Paths: "foxtrot", OrgId: org2ID},
	}
	tests := [...]struct {
		name    string
		OrgId   uuid.UUID
		folders []folder.Folder
		want    []folder.Folder
	}{
		{
			name:    "Get folders for org1",
			OrgId:   uuid.FromStringOrNil("00000000-0000-0000-0000-000000000001"), // Example UUID for org1
			folders: folderData,
			want: []folder.Folder{
				{Name: "alpha", Paths: "alpha", OrgId: org1ID},
				{Name: "bravo", Paths: "alpha.bravo", OrgId: org1ID},
				{Name: "charlie", Paths: "alpha.bravo.charlie", OrgId: org1ID},
				{Name: "delta", Paths: "alpha.delta", OrgId: org1ID},
				{Name: "echo", Paths: "echo", OrgId: org1ID},
			},
		},
		{
			name:    "Get folders for org2",
			OrgId:   uuid.FromStringOrNil("00000000-0000-0000-0000-000000000002"), // Example UUID for org2
			folders: folderData,
			want: []folder.Folder{
				{Name: "foxtrot", Paths: "foxtrot", OrgId: org2ID},
			},
		},
		{
			name:    "Get folders for non-existent org",
			OrgId:   uuid.FromStringOrNil("00000000-0000-0000-0000-000000000003"), // Non-existent org
			folders: folderData,
			want:    []folder.Folder{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			get := f.GetFoldersByOrgID(tt.OrgId)
			if len(get) != len(tt.want) {
				t.Errorf("got %d folders, want %d", len(get), len(tt.want))
			}
			for i := range get {
				if get[i] != tt.want[i] {
					t.Errorf("got %v, want %v", get[i], tt.want[i])
				}
			}
		})
	}
}

func Test_folder_GetAllChildFolders(t *testing.T) {
	t.Parallel()

	// Sample folder data with OrgId as UUIDs
	org1ID := uuid.FromStringOrNil("00000000-0000-0000-0000-000000000001") // Example UUID for org1
	org2ID := uuid.FromStringOrNil("00000000-0000-0000-0000-000000000002") // Example UUID for org2

	folderData := []folder.Folder{
		{OrgId: org1ID, Paths: "alpha", Name: "alpha"},
		{OrgId: org1ID, Paths: "alpha.bravo", Name: "bravo"},
		{OrgId: org1ID, Paths: "alpha.bravo.charlie", Name: "charlie"},
		{OrgId: org1ID, Paths: "alpha.delta", Name: "delta"},
		{OrgId: org1ID, Paths: "echo", Name: "echo"},
		{OrgId: org2ID, Paths: "foxtrot", Name: "foxtrot"},
	}

	tests := [...]struct {
		name       string
		OrgId      uuid.UUID
		folderName string
		folders    []folder.Folder
		want       []folder.Folder
		wantErr    bool
	}{
		{
			name:       "Get all child folders of alpha",
			OrgId:      org1ID,
			folderName: "alpha",
			folders:    folderData,
			want: []folder.Folder{
				{OrgId: org1ID, Paths: "alpha.bravo", Name: "bravo"},
				{OrgId: org1ID, Paths: "alpha.bravo.charlie", Name: "charlie"},
				{OrgId: org1ID, Paths: "alpha.delta", Name: "delta"},
			},
			wantErr: false,
		},
		{
			name:       "Get all child folders of bravo",
			OrgId:      org1ID,
			folderName: "bravo",
			folders:    folderData,
			want: []folder.Folder{
				{OrgId: org1ID, Paths: "alpha.bravo.charlie", Name: "charlie"},
			},
			wantErr: false,
		},
		{
			name:       "Get all child folders of charlie",
			OrgId:      org1ID,
			folderName: "charlie",
			folders:    folderData,
			want:       []folder.Folder{},
			wantErr:    false,
		},
		{
			name:       "Get all child folders of echo",
			OrgId:      org1ID,
			folderName: "echo",
			folders:    folderData,
			want:       []folder.Folder{},
			wantErr:    false,
		},
		{
			name:       "Get all child folders of an invalid folder",
			OrgId:      org1ID,
			folderName: "invalid_folder",
			folders:    folderData,
			want:       []folder.Folder{},
			wantErr:    true,
		},
		{
			name:       "Get all child folders of foxtrot from org1",
			OrgId:      org1ID,
			folderName: "foxtrot",
			folders:    folderData,
			want:       []folder.Folder{},
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			get, error := f.GetAllChildFolders(tt.OrgId, tt.folderName)

			if (error != nil) != tt.wantErr {
				t.Errorf("got error %v, wantErr %v", error, tt.wantErr)
				return
			}

			if len(get) != len(tt.want) {
				t.Errorf("got %d child folders, want %d", len(get), len(tt.want))
			}
			for i := range get {
				if get[i] != tt.want[i] {
					t.Errorf("got %v, want %v", get[i], tt.want[i])
				}
			}
		})
	}
}
