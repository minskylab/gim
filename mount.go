package gim

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strings"
)

func (g *Gim) fileInBlackList( filename string, extra ...string) bool {
	blackList := []string{g.Workspace.Config.GimFolder, "", ".", ".cache", ".git", "node-modules"}
	for _, e := range extra {
		blackList = append(blackList, e)
	}

	for _, black := range blackList {
		if filename == black {
			return true
		}
	}
	return false
}


func (g *Gim) mountDist(location string, routes routesTree) (map[string]string, error) {
	deep := len(strings.Split(location, "/"))
	if deep == 1 {
		_ = g.removeContents(g.Workspace.scriptsPublicFolder())
		_ = os.MkdirAll(g.Workspace.scriptsPublicFolder(), os.ModePerm)
	}

	tableRoutes := make(map[string]string)

	for p, route := range routes {
		if len(route) > 0 {
			newRoutes, err := g.mountDist(path.Join(location, p), route)
			if err != nil {
				return nil, err
			}
			for k, v := range newRoutes {
				tableRoutes[k] = v
			}
		} else {
			fixedLocation := strings.Replace(location, g.Workspace.Config.GimFolder, g.Workspace.Config.PagesFolder, 1)
			filepath := path.Join(g.Workspace.Path, fixedLocation, p)
			hashName := g.generateLittleHash(8) + ".js"

			file, err := ioutil.TempFile(os.TempDir(), "mint")
			if err != nil {
				return nil, err
			}

			data, err := ioutil.ReadFile(filepath)
			if err != nil {
				return nil, err
			}

			if err = ioutil.WriteFile(file.Name(), data, 0644); err != nil {
				return nil, err
			}


			response, err := exec.Command(g.Workspace.Config.ImbacCommand, file.Name()).Output()
			if err != nil {
				fmt.Println(string(response))
				return nil, err
			}

			newData, err := ioutil.ReadAll(file)
			if err != nil {
				return nil, err
			}

			targetFile := path.Join(g.Workspace.Path, g.Workspace.Config.GimFolder, g.Workspace.Config.DistPublicFolder, hashName)

			if err = ioutil.WriteFile(targetFile, newData, 0644); err != nil {
				return nil, err
			}

			prefixToCut := path.Join(g.Workspace.Path, g.Workspace.Config.PagesFolder)
			virtualPath := strings.Replace(filepath, prefixToCut, "", 1)

			prefixToCut = path.Join(g.Workspace.Path, g.Workspace.Config.GimFolder)
			realPath := strings.Replace(targetFile, prefixToCut, "", 1)

			tableRoutes[virtualPath] = realPath
		}

	}

	return tableRoutes, nil
}

func (g *Gim) remountDist(location string, tableRoutes map[string]string) error {
	deep := len(strings.Split(location, "/"))
	if deep == 1 {
		_ = g.removeContents(g.Workspace.scriptsPublicFolder())
		_ = os.MkdirAll(g.Workspace.scriptsPublicFolder(), os.ModePerm)
	}

	for virtualPath, realPath := range tableRoutes {
		filepath := path.Join(g.Workspace.pagesFolder(), virtualPath)
		targetFile := path.Join(g.Workspace.distFolder(), realPath)

		file, err := ioutil.TempFile(os.TempDir(), "mint")
		if err != nil {
			return err
		}

		data, err := ioutil.ReadFile(filepath)
		if err != nil {
			return err
		}

		if err = ioutil.WriteFile(file.Name(), data, 0644); err != nil {
			return err
		}

		response, err := exec.Command(g.Workspace.Config.ImbacCommand, file.Name()).Output()
		if err != nil {
			fmt.Println(string(response))
			return err
		}

		newData, err := ioutil.ReadAll(file)
		if err != nil {
			return err
		}

		if err = ioutil.WriteFile(targetFile, newData, 0644); err != nil {
			return err
		}

	}

	return nil
}


func (g *Gim) mirrorProjectToContent() error {
	files, err := ioutil.ReadDir(g.Workspace.Path)
	if err != nil {
		return err
	}

	for _, file := range files {
		if g.fileInBlackList(file.Name(), g.Workspace.Config.PagesFolder, g.Workspace.Config.PublicFolder) {
			continue
		}

		from := path.Join(g.Workspace.Path, file.Name())
		to := path.Join(g.Workspace.scriptsPublicFolder(), file.Name())
		if err := exec.Command("cp", "-r", from, to).Run(); err != nil {
			return err
		}
	}

	resp, err := exec.Command(g.Workspace.Config.ImbacCommand, g.Workspace.scriptsPublicFolder()).Output()
	if err != nil {
		fmt.Println(string(resp))
		return err
	}

	return nil
}
