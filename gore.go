package main

import (
	"fmt"
	"gore/args"
	"gore/config"
	"gore/db"
	"gore/model"
	"gore/s3"
	"gore/util"
)

func main() {
	args.GetAppName()
	
	util.SetReleaseDir()
	util.AppExistsInReleaseDirOrExit()

	util.PrintGoreLogo()

	config.LoadGoreConfig()
	db.OpenDb()

	release, exists := db.GetLatestVersionOfApp(args.App)

	if !exists {
		fmt.Println()
		fmt.Println("This app has not been released yet.")
		fmt.Printf("This will be the v1.0.0 release for '%s'.", args.App)

		firstRelease := model.Release{App: args.App, Major: "1", Minor: "0", Patch: "0"}
		ReleaseApp(firstRelease)

		return
	}

	newRelease := model.GetNewRelease(release)
	ReleaseApp(newRelease)
}

func ReleaseApp(release model.Release) {
	s3.ConfirmUpload(release)
	s3.UploadReleaseToS3(release)
	db.SetAppVersion(release)
	fmt.Printf("\nSuccessfully uploaded %s.\n", util.Green(release.QualifiedName()))
}
