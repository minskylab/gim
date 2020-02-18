package gim

import (
	"os"
	"path/filepath"
	"strings"
)

type routesTree map[string]routesTree

type routeDescription struct {
	path      string
	parameter string
}

func (g *Gim) extractRouteTree() (routesTree, error) {
	routesRoot := routesTree{}

	deep := 0
	if err := filepath.Walk(g.Workspace.pagesFolder(), func(p string, info os.FileInfo, err error) error {
		relativePath := g.cutPath(g.Workspace.pagesFolder(), p)

		parts := strings.Split(relativePath, "/")

		if len(parts) == deep+1 {
			deep++
		}

		if !info.IsDir() {
			var current *routesTree
			current = &routesRoot
			for _, part := range parts {
				if (*current)[part] == nil {
					(*current)[part] = routesTree{}
				}
				var temp = (*current)[part]

				current = &temp
			}
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return routesRoot, nil
}

func (g *Gim) describeRoute(filename string) routeDescription {
	chunks := strings.Split(filename, ".")
	if len(chunks) == 0 {
		return routeDescription{}
	}

	name := chunks[0]
	// TODO: You can validate if the extension is valid
	description := routeDescription{
		path:      name,
		parameter: "",
	}

	if name == "index" {
		description.path = ""
	}

	if strings.HasPrefix(name, "[") && strings.HasSuffix(name, "]") {
		parameterName := strings.Trim(name, "[]")
		description.path = ":" + parameterName
		description.parameter = parameterName
	}

	return description
}
