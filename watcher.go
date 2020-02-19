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
						fmt.Println(err)
					}

					if err := g.updateWatcherScope(); err != nil {
						fmt.Println(err)
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

	g.watcher = watcher
	return g.updateWatcherScope()
}

func (g *Gim) updateWatcherScope() error {
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
			err = g.watcher.Add(p)
			if err != nil {
				return err
			}
		}
		return nil
	})
}