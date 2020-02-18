package gim

import "path"

type Workspace struct {
	Path string
	Config *Config
}

func NewWorkspace(path string, conf *Config) *Workspace {
	return &Workspace{
		Path:   path,
		Config: conf,
	}
}

func (w *Workspace) pagesFolder() string {
	return path.Join(w.Path, w.Config.PagesFolder)
}

func (w *Workspace) publicFolder() string {
	return path.Join(w.Path, w.Config.PublicFolder)
}

func (w *Workspace) distFolder() string {
	return path.Join(w.Path, w.Config.GimFolder)
}


func (w *Workspace) scriptsPublicFolder() string {
	return path.Join(w.Path, w.Config.GimFolder, w.Config.DistPublicFolder)
}

func (w *Workspace) scriptsBrowserFolder() string {
	return path.Join(w.Path, w.Config.GimFolder, w.Config.DistBrowserFolder)
}

func (w *Workspace) mainTemplate() string {
	return path.Join(w.Path, w.Config.GimFolder, w.Config.TemplateHTMLName)
}