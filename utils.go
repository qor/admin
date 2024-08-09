package admin

import (
	"html/template"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/qor/assetfs"
	"github.com/qor/qor"
	"github.com/qor/qor/utils"
	"github.com/qor/roles"
)

var (
	globalViewPaths []string
	globalAssetFSes []assetfs.Interface
	goModDeps       []string
	goModPrefix     = "pkg/mod"
)

// HasPermissioner has permission interface
type HasPermissioner interface {
	HasPermission(roles.PermissionMode, *qor.Context) bool
}

// ResourceNamer is an interface for models that defined method `ResourceName`
type ResourceNamer interface {
	ResourceName() string
}

// I18n define admin's i18n interface
type I18n interface {
	Scope(scope string) I18n
	Default(value string) I18n
	T(locale string, key string, args ...interface{}) template.HTML
}

// RegisterViewPath register view path for all assetfs
func RegisterViewPath(pth string) {
	globalViewPaths = append(globalViewPaths, pth)
	var err error
	for _, assetFS := range globalAssetFSes {
		if err = assetFS.RegisterPath(filepath.Join(utils.AppRoot, "vendor", pth)); err != nil {
			for _, gopath := range utils.GOPATH() {
				startDir, _ := os.Getwd()
				if modPath := findGoModFile(startDir); modPath != "" {
					if err = assetFS.RegisterPath(filepath.Join(gopath, getDepVersionFromMod(modPath, pth))); err == nil {
						break
					}
				}

				if err = assetFS.RegisterPath(filepath.Join(gopath, "src", pth)); err == nil {
					break
				}
			}
		}
	}
	if err != nil {
		log.Printf("RegisterViewPathError: %s %s!", pth, err.Error())
	}
}

func equal(a, b interface{}) bool {
	return reflect.DeepEqual(a, b)
}

func getDepVersionFromMod(modPath, pth string) string {
	if len(goModDeps) == 0 {
		if cont, err := os.ReadFile(modPath); err == nil {
			goModDeps = strings.Split(string(cont), "\n")
		}
	}

	for _, val := range goModDeps {
		if txt := strings.Trim(val, "\t\r"); strings.HasPrefix(txt, pth) {
			return filepath.Join(goModPrefix, pth+"@"+strings.Split(txt, " ")[1])
		}
	}

	if strings.LastIndex(pth, "/") == -1 {
		return pth
	}

	return getDepVersionFromMod(modPath, pth[:strings.LastIndex(pth, "/")]) + pth[strings.LastIndex(pth, "/"):]
}

func findGoModFile(startDir string) string {
	currentDir := startDir

	for {
		goModPath := filepath.Join(currentDir, "go.mod")
		if _, err := os.Stat(goModPath); err == nil {
			return goModPath
		}

		parentDir := filepath.Dir(currentDir)
		if parentDir == currentDir {
			break
		}
		currentDir = parentDir
	}

	return ""
}
