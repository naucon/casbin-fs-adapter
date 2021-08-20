package fsadapter

import (
	"github.com/casbin/casbin/v2/model"
	"io/fs"
)

func NewModel(fsys fs.FS, filePath string) (model.Model, error) {
	b, err := fs.ReadFile(fsys, filePath)
	if err != nil {
		return nil, err
	}

	return model.NewModelFromString(string(b))
}
