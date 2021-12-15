package util

import (
	"fmt"
	"gore/args"
	"io/ioutil"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var ReleaseDir string

func SetReleaseDir() {
	ReleaseDir = HomeDir() + "/release/"
}

func AppExistsInReleaseDirOrExit() {
	fileExists := DoesFileExist(ReleaseDir, args.App)
	if !fileExists {
		fmt.Printf("File '%s' does not exist in ~/release/\n", args.App)
		os.Exit(0)
	}
}

func DoesFileExist(dir string, name string) bool {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatalf("Failed to list release dir: %+v", err)
	}
	for _, file := range files {
		if file.Name() == name {
			return true
		}
	}
	return false
}

func IsIn(target string, arr []string) bool {
	for _, str := range arr {
		if target == str {
			return true
		}
	}
	return false
}

func HomeDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Could not get user's home directory: %+v", err)
	}
	return home
}
