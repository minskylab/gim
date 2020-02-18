package gim

import (
	"time"

	"github.com/briandowns/spinner"

)

func (g *Gim) regenerateScripts(tableRoutes map[string]string) error {
	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)

	g.Printer.ShowLoading("Mounting .mint dir")
	err := g.remountDist(g.Workspace.Config.MintFolder, tableRoutes)
	if err != nil {
		return err
	}
	g.Printer.HideLoading()

	g.Printer.ShowLoading("Bootstrap .mint dir")
	s.Start()
	if err := g.bootstrapDist(); err != nil {
		return err
	}
	g.Printer.HideLoading()

	g.Printer.ShowLoading("Mirroring extra project content .mint dir")
	s.Start()
	if err := g.mirrorProjectToContent(); err != nil {
		return err
	}
	g.Printer.HideLoading()

	g.Printer.ShowLoading("Building browser scripts")
	s.Start()
	if err := g.Builder.BuildBrowserScripts(g.Workspace, true); err != nil {
		return err
	}
	g.Printer.HideLoading()

	return nil
}
