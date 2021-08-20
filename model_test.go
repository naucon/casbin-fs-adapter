package fsadapter

import (
	"github.com/casbin/casbin/v2/model"
	"github.com/stretchr/testify/assert"
	"io/fs"
	"os"
	"testing"
)

func TestCasbinFsAdapter_NewModel(t *testing.T) {
	t.Run("TestCasbinFsAdapter_NewModel_ShouldReturnModel", func(t *testing.T) {
		fsys := os.DirFS("examples/config/")
		actualModel, err := NewModel(fsys, "model.conf")
		if err != nil {
			assert.FailNow(t, err.Error())
		}
		assert.IsType(t, model.Model{}, actualModel)
	})

	t.Run("TestCasbinFsAdapter_NewModel_MissingShouldReturnError", func(t *testing.T) {
		fsys := os.DirFS("examples/config/")
		_, err := NewModel(fsys, "missing.conf")
		assert.Error(t, err)
		assert.ErrorIs(t, err, fs.ErrNotExist)
	})
}
