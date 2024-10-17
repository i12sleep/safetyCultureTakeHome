package folder_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
)

func Test_folder_MoveFolder(t *testing.T) {
	org1 := uuid.Must(uuid.NewV4())
	org2 := uuid.Must(uuid.NewV4())

	folders := []folder.Folder{
		{Name: "alpha", Paths: "alpha", OrgId: org1},
		{Name: "bravo", Paths: "alpha.bravo", OrgId: org1},
		{Name: "charlie", Paths: "alpha.bravo.charlie", OrgId: org1},
		{Name: "delta", Paths: "alpha.delta", OrgId: org1},
		{Name: "echo", Paths: "alpha.delta.echo", OrgId: org1},
		{Name: "foxtrot", Paths: "foxtrot", OrgId: org2},
		{Name: "golf", Paths: "golf", OrgId: org1},
	}

	tests := []struct {
		name      string
		src       string
		dst       string
		want      []folder.Folder
		wantErr   bool
		errString string
	}{
		{
			name: "Move bravo to delta",
			src:  "bravo",
			dst:  "delta",
			want: []folder.Folder{
				{Name: "alpha", Paths: "alpha", OrgId: org1},
				{Name: "bravo", Paths: "alpha.delta.bravo", OrgId: org1},
				{Name: "charlie", Paths: "alpha.delta.bravo.charlie", OrgId: org1},
				{Name: "delta", Paths: "alpha.delta", OrgId: org1},
				{Name: "echo", Paths: "alpha.delta.echo", OrgId: org1},
				{Name: "foxtrot", Paths: "foxtrot", OrgId: org2},
				{Name: "golf", Paths: "golf", OrgId: org1},
			},
			wantErr: false,
		},
		{
			name: "Move bravo to golf",
			src:  "bravo",
			dst:  "golf",
			want: []folder.Folder{
				{Name: "alpha", Paths: "alpha", OrgId: org1},
				{Name: "bravo", Paths: "golf.bravo", OrgId: org1},
				{Name: "charlie", Paths: "golf.bravo.charlie", OrgId: org1},
				{Name: "delta", Paths: "alpha.delta", OrgId: org1},
				{Name: "echo", Paths: "alpha.delta.echo", OrgId: org1},
				{Name: "foxtrot", Paths: "foxtrot", OrgId: org2},
				{Name: "golf", Paths: "golf", OrgId: org1},
			},
			wantErr: false,
		},
		{
			name:      "Move bravo to charlie (invalid)",
			src:       "bravo",
			dst:       "charlie",
			wantErr:   true,
			errString: "error: Cannot move a folder to a child of itself",
		},
		{
			name:      "Move bravo to bravo (invalid)",
			src:       "bravo",
			dst:       "bravo",
			wantErr:   true,
			errString: "error: Cannot move a folder to itself",
		},
		{
			name:      "Move bravo to foxtrot (different org)",
			src:       "bravo",
			dst:       "foxtrot",
			wantErr:   true,
			errString: "error: Cannot move a folder to a different organization",
		},
		{
			name:      "Move invalid_folder to delta (source folder does not exist)",
			src:       "invalid_folder",
			dst:       "delta",
			wantErr:   true,
			errString: "error: Source folder does not exist",
		},
		{
			name:      "Move bravo to invalid_folder (destination folder does not exist)",
			src:       "bravo",
			dst:       "invalid_folder",
			wantErr:   true,
			errString: "error: Destination folder does not exist",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(folders)

			got, err := f.MoveFolder(tt.src, tt.dst)
			if (err != nil) != tt.wantErr {
				t.Errorf("MoveFolder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr && err != nil && err.Error() != tt.errString {
				t.Errorf("MoveFolder() error = %v, expected error %v", err.Error(), tt.errString)
				return
			}

			if !tt.wantErr {
				if !compareFolders(got, tt.want) {
					t.Errorf("MoveFolder() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func compareFolders(a, b []folder.Folder) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
