package folder

import (
	"errors"
	"strings"

	"github.com/gofrs/uuid"
)

func GetAllFolders() []Folder {
	return GetSampleData()
}

func (f *driver) GetFoldersByOrgID(orgID uuid.UUID) []Folder {
	folders := f.folders

	res := []Folder{}
	for _, f := range folders {
		if f.OrgId == orgID {
			res = append(res, f)
		}
	}

	return res

}

func (f *driver) GetAllChildFolders(orgID uuid.UUID, name string) ([]Folder, error) {
	var parent *Folder
	folders := f.GetFoldersByOrgID(orgID)

	if folders == nil {
		return nil, errors.New("Error: Folder does not exist")
	}

	for _, folder := range folders {
		if folder.Name == name {
			parent = &folder
		}
	}

	if parent == nil {
		return nil, errors.New("Error: Folder does not exist in the specified organization")
	}

	var child []Folder
	for _, folder := range folders {
		if strings.HasPrefix(folder.Paths, parent.Paths+".") {
			child = append(child, folder)
		}
	}

	return child, nil
}
