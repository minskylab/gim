package gim


type Builder interface {
	BuildBrowserScripts(workspace *Workspace, format bool) error
}