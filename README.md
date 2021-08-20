# FS Adapter for Casbin

[![Build](https://github.com/naucon/casbin-fs-adapter/actions/workflows/go-ci.yml/badge.svg)](https://github.com/naucon/casbin-fs-adapter/actions/workflows/go-ci.yml)
[![Coverage](https://codecov.io/gh/naucon/casbin-fs-adapter/branch/master/graph/badge.svg?token=YQ5BQU03PE)](https://codecov.io/gh/naucon/casbin-fs-adapter)

This package is a file system adapter for [Casbin](https://github.com/casbin/casbin) v2.
The adapter enables [Casbin](https://github.com/casbin/casbin) to load policies and models from fs.FS interface including support for embed.FS.

**NOTICE:** because fs.FS is readonly, writing operations and AutoSave are not supported.

## Requires

* Go 1.16 or newer

## Installation

install the latest version via go get

```
go get -u github.com/naucon/casbin-fs-adapter
```

## Import package

```
import (
  casbin_fs_adapter "github.com/naucon/casbin-fs-adapter"
)
```

## Usage

### Policy

This Adapter implements the Casbin storage policy adapter interface and can be injected into the Casbin Enforcer.
To inject the storage adapter, create an instance with `casbin_fs_adapter.NewAdapter()`. Pass in the filesystem and file path.

```
	fsys := os.DirFS("config")
	policies := casbin_fs_adapter.NewAdapter(fsys, "policy.csv")
	enforcer, _ := casbin.NewEnforcer("config/casbin_model.conf", policies)
```

### Model

Currently, Casbin has no storage adapter interface for model configuration. However, the model can be created and injected into the casbin Enforcer.
To create a model call `casbin_fs_adapter.NewModel()`, pass in the filesystem and file path.

```
	fsys := os.DirFS("config")
	model, _ := casbin_fs_adapter.NewModel(fsys, "casbin_model.conf")
	policies := casbin_fs_adapter.NewAdapter(fsys, "policy.csv")
	enforcer, _ := casbin.NewEnforcer(model, policies)
```

### Embed

With go:embed we can embed files and directories into application binaries at compile-time.
Embedded directories or multiple files can be access through a variable of type `embed.FS`.
Because `embed.FS` implements the `fs.FS` interface, we can use it in our adapter too.

```
  //go:embed config/model.conf config/policy.csv
  var f embed.FS

  func main() {
      model, _ := casbin_fs_adapter.NewModel(EmbeddedFiles, "config/model.conf")
      policies := casbin_fs_adapter.NewAdapter(EmbeddedFiles, "config/policy.csv")
    	enforcer, _ := casbin.NewEnforcer(model, policies)
  }
```

## Example

````
package main

import (
	"fmt"
	casbin_fs_adapter "github.com/naucon/casbin-fs-adapter"
	"os"

	"github.com/casbin/casbin/v2"
)

func main() {
	fsys := os.DirFS("config")
	model, _ := casbin_fs_adapter.NewModel(fsys, "model.conf")
	policies := casbin_fs_adapter.NewAdapter(fsys, "policy.csv")
	enforcer, _ := casbin.NewEnforcer(model, policies)

	_ = enforcer.LoadPolicy()

	sub := "alice" // user
	obj := "data1" // resource
	act := "read" // operation

	if res, _ := enforcer.Enforce(sub, obj, act); res {
		fmt.Println("permitted")
	} else {
		fmt.Println("rejected")
	}
}
````

## Example with embed

````
package main

import (
	"embed"
	"fmt"
	casbin_fs_adapter "github.com/naucon/casbin-fs-adapter"

	"github.com/casbin/casbin/v2"
)

//go:embed config/model.conf config/policy.csv
var f embed.FS

func main() {
	model, _ := casbin_fs_adapter.NewModel(EmbeddedFiles, "config/model.conf")
	policies := casbin_fs_adapter.NewAdapter(EmbeddedFiles, "config/policy.csv")
	enforcer, _ := casbin.NewEnforcer(model, policies)

	_ = enforcer.LoadPolicy()

	sub := "alice" // user
	obj := "data1" // resource
	act := "read" // operation

	if res, _ := enforcer.Enforce(sub, obj, act); res {
		fmt.Println("permitted")
	} else {
		fmt.Println("rejected")
	}
}
````

## License

This project is licensed under the MIT license. See the [LICENSE](https://github.com/naucon/casbin-fs-adapter/blob/master/LICENSE) file for the full license text.
