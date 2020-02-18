package gim

import (
	"fmt"
	"os/exec"
	"path"
)

type ParcelBuilder struct {
	parcelBinary string
}

func (p *ParcelBuilder) BuildBrowserScripts(wSpace *workspace, format bool) error {
	buildPath := path.Join(wSpace.scriptsPublicFolder(), "*")
	browserTargetPath := path.Join(wSpace.scriptsBrowserFolder())

	if format {
		resp, err := exec.Command("rm", "-rfd", browserTargetPath).Output()
		if err != nil {
			fmt.Println(string(resp))
			return err
		}
	}

	publicBaseURL := "/" + wSpace.Config.DistPublicFolder

	resp, err := exec.Command(p.parcelBinary, "build", buildPath, "-d", browserTargetPath, "--public-url", publicBaseURL).Output()
	if err != nil {
		fmt.Println(string(resp))
		return err
	}

	return nil
}
