package gim

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"
)

func (g *Gim) launchProjectWatcher(tableRoutes map[string]string) error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Write == fsnotify.Write {
					if err := g.regenerateScripts(tableRoutes); err != nil {
						panic(err)
					}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				fmt.Println("error:", err)
			}
		}
	}()

	return filepath.Walk(g.Workspace.Path, func(p string, info os.FileInfo, err error) error {
		rootLevel1 := ""
		// log.Println(p, strings.HasPrefix(p, strings.Trim(g.Workspace.Path, "./")))
		if strings.HasPrefix(p, strings.Trim(g.Workspace.Path, "./")) {
			chunks := strings.Split(p, "/")
			if len(chunks) > 1 {
				rootLevel1 = chunks[0]
			}
		}
		// log.Println("rootLevel1", rootLevel1)
		if !g.fileInBlackList(rootLevel1) {
			err = watcher.Add(p)
			if err != nil {
				return err
			}
		}
		return nil
	})
}