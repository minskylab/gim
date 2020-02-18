package gim

type Router interface  {
	MountStatics(workspace *Workspace) error
	DefineRoutes(workspace *Workspace, tableRoutes map[string]string, parent string, routes routesTree) error
	Run(port int) error
}


