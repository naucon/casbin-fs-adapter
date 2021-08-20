# FS Adapter for casbin

This package is a FileSystem adapter for [Casbin](https://github.com/casbin/casbin) v2.
The adapter enables [Casbin](https://github.com/casbin/casbin) to load policies and models from FileSystem interface including embed.FS.

**NOTICE:** Adapter is readonly, writing or AutoSave is not supported.

## Installation

install the latest version via go get

  go get -u github.com/naucon/casbin-fs-adapter

## Usage

Here is a basic example for using this package:

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
	model, _ := casbin_fs_adapter.NewModel(fsys, "casbin_model.conf")
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

## Example with embed (go1.16+)

````
package main

import (
	"embed"
	"fmt"
	casbin_fs_adapter "github.com/naucon/casbin-fs-adapter"

	"github.com/casbin/casbin/v2"
)

//go:embed config/casbin_model.conf config/policy.csv
var f embed.FS

func main() {
	model, _ := casbin_fs_adapter.NewModel(EmbeddedFiles, "config/casbin_model.conf")
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

This project is licensed under the MIT license. See the LICENSE file for the full license text.
