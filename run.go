package gim

import (
	"fmt"
	"math/rand"
	"time"
)

func (g *Gim) init() {
	rand.Seed(time.Now().Unix())
}


func (g *Gim) Run() error {
	g.Printer.ShowLoading("Walking on pages")
	routes, err := g.extractRouteTree()
	if err != nil {
		return err
	}
	g.Printer.HideLoading()

	g.Printer.ShowLoading(" Mounting .mint dir")
	tableRoutes, err := g.mountDist(g.Workspace.Config.GimFolder, routes)
	if err != nil {
		return err
	}
	g.Printer.HideLoading()

	if err = g.regenerateScripts(tableRoutes); err != nil {
		return err
	}

	if err := g.Router.MountStatics(g.Workspace); err != nil {
		return err
	}

	if err :=  g.Router.DefineRoutes(g.Workspace, tableRoutes, "", routes); err != nil {
		return err
	}

	if err := g.launchProjectWatcher(tableRoutes); err != nil {
		return err
	}

	fmt.Println("Server Ready, listen at :8080")
	return g.Router.Run(8080)
}
