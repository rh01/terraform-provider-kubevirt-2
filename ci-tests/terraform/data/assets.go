<<<<<<< HEAD
//go:build !release
// +build !release

=======
// +build !release
>>>>>>> 0faf8ce (Revert "Upgrade go mod and dependencies")
//go:generate go run assets_generate.go

package data

import (
	"net/http"
	"os"
)

// Assets contains project assets.
var Assets http.FileSystem

func init() {
	dir := os.Getenv("OPENSHIFT_INSTALL_DATA")
	if dir == "" {
		dir = "../terraform/data"
	}
	Assets = http.Dir(dir)
}
