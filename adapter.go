package fsadapter

import (
	"bufio"
	"errors"
	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
	"io/fs"
	"strings"
)

type Adapter struct {
	fsys     fs.FS
	filePath string
}

func NewAdapter(fsys fs.FS, filePath string) *Adapter {
	return &Adapter{fsys, filePath}
}

func (a *Adapter) LoadPolicy(model model.Model) error {
	if a.filePath == "" {
		return errors.New(errInvalidFilePath)
	}

	return a.loadPolicyFile(model, persist.LoadPolicyLine)
}

func (a *Adapter) loadPolicyFile(model model.Model, handler func(string, model.Model)) error {
	f, err := a.fsys.Open(a.filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		handler(line, model)
	}
	return scanner.Err()
}

func (a *Adapter) SavePolicy(model model.Model) error {
	return errors.New(errNotImplemented)
}

func (a *Adapter) AddPolicy(sec string, ptype string, rule []string) error {
	return errors.New(errNotImplemented)
}

func (a *Adapter) RemovePolicy(sec string, ptype string, rule []string) error {
	return errors.New(errNotImplemented)
}

func (a *Adapter) RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error {
	return errors.New(errNotImplemented)
}
