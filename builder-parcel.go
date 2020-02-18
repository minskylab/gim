package gim

import (
	"fmt"
	"os/exec"
	"path"
)

type ParcelBuilder struct {
	parcelBinary string
}

func NewParcelBuilder(parcelBinaryLocation ...string) *ParcelBuilder {
	bin := "parcel"
	if len(parcelBinaryLocation) > 0 {
		bin = parcelBinaryLocation[0]
	}

	return &ParcelBuilder{parcelBinary: bin}
}

func (p *ParcelBuilder) BuildBrowserScripts(wSpace *Workspace, format bool) error {
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
