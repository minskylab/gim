package gim

func (g *Gim) regenerateScripts(tableRoutes map[string]string) error {

	g.Printer.ShowLoading("Mounting .mint dir")
	if err := g.remountDist(g.Workspace.Config.GimFolder, tableRoutes); err != nil {
		return err
	}
	g.Printer.HideLoading()

	g.Printer.ShowLoading("Bootstrap .mint dir")
	if err := g.bootstrapDist(); err != nil {
		return err
	}
	g.Printer.HideLoading()

	g.Printer.ShowLoading("Mirroring extra project content .mint dir")
	if err := g.mirrorProjectToContent(); err != nil {
		return err
	}
	g.Printer.HideLoading()

	g.Printer.ShowLoading("Building browser scripts")
	if err := g.Builder.BuildBrowserScripts(g.Workspace, true); err != nil {
		return err
	}
	g.Printer.HideLoading()

	return nil
}
