package load

import (
	"golang.org/x/tools/go/packages"
)

const loadMode = packages.NeedName | packages.NeedFiles | packages.NeedCompiledGoFiles |
	packages.NeedImports | packages.NeedTypes | packages.NeedTypesSizes | packages.NeedSyntax |
	packages.NeedExportsFile

var (
	config = &packages.Config{
		Mode: loadMode,
	}
)

func LoadFile(filename string) ([]*packages.Package, error) {
	pkgs, err := packages.Load(config, filename)
	if err != nil {
		return nil, err
	}

	return pkgs, nil
}
