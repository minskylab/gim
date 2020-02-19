package gim

import (
	"fmt"
	"io/ioutil"
	"path"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/robertkrimen/otto"
)

const configFilename = "_gimConfig"

type GinRouter struct {
	engine *gin.Engine
}

func NewGinRouter() *GinRouter {
	gin.SetMode(gin.TestMode)
	return &GinRouter{
		engine: gin.New(),
	}
}

func (r *GinRouter) fileInBlackList(workspace *Workspace, filename string, extra ...string) bool {
	blackList := []string{workspace.Config.GimFolder, "", ".", ".cache", ".git", "node_modules"}
	for _, e := range extra {
		blackList = append(blackList, e)
	}

	for _, black := range blackList {
		if filename == black {
			return true
		}
	}
	return false
}


func (r *GinRouter) describeRoute(filename string) routeDescription {
	chunks := strings.Split(filename, ".")
	if len(chunks) == 0 {
		return routeDescription{}
	}

	name := chunks[0]
	// TODO: You can validate if the extension is valid
	description := routeDescription{
		path:      name,
		parameter: "",
	}

	if name == "index" {
		description.path = ""
	}

	if strings.HasPrefix(name, "[") && strings.HasSuffix(name, "]") {
		parameterName := strings.Trim(name, "[]")
		description.path = ":" + parameterName
		description.parameter = parameterName
	}

	return description
}


func (r *GinRouter) MountStatics(wSpace *Workspace) error {
	r.engine.Static("/"+ wSpace.Config.DistPublicFolder, wSpace.scriptsBrowserFolder())
	r.engine.Static("/"+ wSpace.Config.PublicFolder, wSpace.publicFolder())
	r.engine.LoadHTMLFiles(wSpace.mainTemplate())

	files, err := ioutil.ReadDir(wSpace.Path)
	if err != nil {
		return err
	}

	for _, file := range files {
		if r.fileInBlackList(wSpace, file.Name(), wSpace.Config.PagesFolder, wSpace.Config.PublicFolder) {
			continue
		}

		absPath := path.Join(wSpace.Path, wSpace.Config.GimFolder, wSpace.Config.DistBrowserFolder, file.Name())
		r.engine.Static("/"+file.Name(), absPath)
	}

	return nil
}

func (r *GinRouter) DefineRoutes(wSpace *Workspace, tableRoutes map[string]string, parent string, routes routesTree) error {
	for p, route := range routes {
		if len(route) > 0 {
			if err := r.DefineRoutes(wSpace, tableRoutes, strings.Join([]string{parent, p}, "/"), route); err != nil {
				return err
			}
		} else {
			routeRest := strings.Join([]string{parent, r.describeRoute(p).path}, "/")
			source, ok := tableRoutes[parent+"/"+p]
			if !ok {
				source = "xxx.js"
			}

			r.engine.GET(routeRest, func(c *gin.Context) {
				targetFile := path.Join(wSpace.distFolder(), source)
				data, err := ioutil.ReadFile(targetFile)
				if err != nil {
					c.String(500, err.Error())
					return
				}

				params := gin.H{"Source": source}

				configString := regexp.MustCompile(`\w+ `+configFilename+` = {\n*([^}]+)}`).FindString(string(data))
				if strings.HasPrefix(configString, "const") {
					configString = strings.Replace(configString, "const", "var", 1)
				}

				styleRequirements := regexp.MustCompile(`require\(["'].*\.css["']\)`).FindAllString(string(data), -1)

				vm := otto.New()
				if _, err := vm.Run(configString); err != nil {
					c.String(500, err.Error())
					return
				}

				conf, err := vm.Get(configFilename)
				if err != nil {
					c.String(500, err.Error())
					return
				}

				if conf.IsObject() {
					title, err := conf.Object().Get("title")
					if err != nil {
						c.String(500, err.Error())
						return
					}

					if title.IsString() {
						params["Title"] = title.String()
					}

					// metas, err := conf.Object().Get("metas")
					// if err != nil {
					// 	c.String(500, err.Error())
					// 	return
					// }

					params["Metas"] = []map[string]string{}


					if len(styleRequirements) > 0 {
						if params["Stylesheets"] == nil {
							params["Stylesheets"] = []map[string]string{}
						}
						params["Stylesheets"] = append(params["Stylesheets"].([]map[string]string), map[string]string{
							"Source": strings.Replace(source, ".js", ".css", 1),
						})
					}

				}

				c.HTML(200, "template.html", params)
			})
		}
	}

	return nil
}

func (r *GinRouter) Run(port int) error {
	return r.engine.Run(fmt.Sprintf(":%d", port))
}