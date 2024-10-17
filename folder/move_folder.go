package folder

import (
	"errors"
	"strings"
)

func (f *driver) MoveFolder(name string, dst string) ([]Folder, error) {
	if name == dst {
		return nil, errors.New("error: Cannot move a folder to itself")
	}

	var move *Folder
	for i, folder := range f.folders {
		if folder.Name == name {
			move = &f.folders[i]
			break
		}
	}

	if move == nil {
		return nil, errors.New("error: Source folder does not exist")
	}

	var dest *Folder
	for i, folder := range f.folders {
		if folder.Name == dst {
			dest = &f.folders[i]
			break
		}
	}

	if dest == nil {
		return nil, errors.New("error: Destination folder does not exist")
	}

	if dest.OrgId != move.OrgId {
		return nil, errors.New("error: Cannot move a folder to a different organization")
	}

	children, err := f.GetAllChildFolders(move.OrgId, move.Name)
	if err != nil {
		return nil, err
	}

	if strings.HasPrefix(dest.Paths, move.Paths) {
		return nil, errors.New("error: Cannot move a folder to a child of itself")
	}

	oldPath := move.Paths
	newPath := dest.Paths + "." + move.Name

	move.Paths = newPath

	for _, child := range children {
		for i, folder := range f.folders {
			if folder.Name == child.Name {
				newChildPath := strings.Replace(f.folders[i].Paths, oldPath, newPath, 1)
				f.folders[i].Paths = newChildPath
			}
		}
	}
	return f.folders, nil
}
