package model

import "fmt"

type Release struct {
	App   string `db:"app"`
	Major string `db:"major"`
	Minor string `db:"minor"`
	Patch string `db:"patch"`
}

func (release Release) QualifiedName() string {
	return release.App + "-" + release.SemVer()
}

func (release Release) SemVer() string {
	return fmt.Sprintf("v%s.%s.%s", release.Major, release.Minor, release.Patch)
}
