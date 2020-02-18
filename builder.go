package gim


type Builder interface {
	BuildBrowserScripts(workspace *workspace, format bool) error
}