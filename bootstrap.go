package gim

import (
	"io/ioutil"
	"path"
)

const defaultTemplate = `<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	{{if .Title}}
	<title>{{.Title}}</title>
	{{else}}
	<title>Mint</title>
	{{end}}
	{{if .Stylesheets}}
	{{range .Stylesheets}}
	<link rel="stylesheet" type="text/css" href="{{.Source}}"/>
	{{end}}
	{{end}}
	{{if .Metas}}
	{{range .Metas}}
	<meta name="{{.Name}}" content="{{.Content}}">
	{{end}}
	{{end}}
</head>
<body>
	<script src="{{.Source}}"></script>
</body>
</html>
`

func (g *Gim) bootstrapDist(templateFile ...string) error {
	filename := path.Join(g.Workspace.Path, g.Workspace.Config.GimFolder, g.Workspace.Config.TemplateHTMLName)
	if err := ioutil.WriteFile(filename, []byte(defaultTemplate), 0644); err != nil {
		return err
	}

	return nil
}

