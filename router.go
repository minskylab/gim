package gim

type Router interface  {
	MountStatics(workspace *workspace) error
	DefineRoutes(workspace *workspace, tableRoutes map[string]string, parent string, routes routesTree) error
	Run(port int) error
}


