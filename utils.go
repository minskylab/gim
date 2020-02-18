package gim

import (
	"math/rand"
	"os"
	"path"
	"path/filepath"
	"strings"
)

const alphabet = "abcdefghijklmnopqrst0123456789"

func (g *Gim) cutPath(father string, value string) string {
	cutestPath := strings.Replace(value, father, "", 1)
	return path.Join(".", cutestPath)
}

func (g *Gim) normalize(routes routesTree) routesTree {
	vRoutes := make(routesTree)
	for k, r := range routes {
		if len(r) != 0 {
			norm := g.normalize(r)
			vRoutes[k] = norm
		} else {
			vRoutes[k] = routesTree{}
		}
	}

	return vRoutes
}

func (g *Gim) generateLittleHash(size int) string {
	hash := ""
	for i := 0; i < size; i++ {
		hash += string(alphabet[rand.Intn(len(alphabet))])
	}

	return hash
}

func (g *Gim) removeContents(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return nil
}