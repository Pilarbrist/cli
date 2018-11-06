package v7pushaction

import (
	"io"

	"code.cloudfoundry.org/cli/actor/sharedaction"
	"code.cloudfoundry.org/cli/actor/v7action"
)

//go:generate counterfeiter . V3Actor

type V7Actor interface {
	CloudControllerAPIVersion() string
	CreateApplicationInSpace(app v7action.Application, spaceGUID string) (v7action.Application, v7action.Warnings, error)
	CreateBitsPackageByApplication(appGUID string) (v7action.Package, v7action.Warnings, error)
	GetApplicationByNameAndSpace(appName string, spaceGUID string) (v7action.Application, v7action.Warnings, error)
	PollPackage(pkg v7action.Package) (v7action.Package, v7action.Warnings, error)
	SetApplicationDroplet(appGUID string, dropletGUID string) (v7action.Warnings, error)
	StageApplicationPackage(pkgGUID string) (v7action.Build, v7action.Warnings, error)
	PollBuild(buildGUID string, appName string) (v7action.Droplet, v7action.Warnings, error)
	UpdateApplication(v7action.Application) (v7action.Application, v7action.Warnings, error)
	UploadBitsPackage(v7action.Package, []sharedaction.Resource, io.Reader, int64) (v7action.Package, v7action.Warnings, error)
}
