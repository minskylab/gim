package gim

import "path"

type workspace struct {
	Path string
	Config *config
}

func newWorkspace(path string, conf *config) *workspace {
	return &workspace{
		Path:   path,
		Config: conf,
	}
}

func (w *workspace) pagesFolder() string {
	return path.Join(w.Path, w.Config.PagesFolder)
}

func (w *workspace) publicFolder() string {
	return path.Join(w.Path, w.Config.PublicFolder)
}

func (w *workspace) distFolder() string {
	return path.Join(w.Path, w.Config.MintFolder)
}


func (w *workspace) scriptsPublicFolder() string {
	return path.Join(w.Path, w.Config.MintFolder, w.Config.DistPublicFolder)
}

func (w *workspace) scriptsBrowserFolder() string {
	return path.Join(w.Path, w.Config.MintFolder, w.Config.DistBrowserFolder)
}

func (w *workspace) mainTemplate() string {
	return path.Join(w.Path, w.Config.MintFolder, w.Config.TemplateHTMLName)
}