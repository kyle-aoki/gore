package s3

import (
	"gore/args"
	"gore/util"
	"log"
	"os"
)

func ReadReleaseFile() []byte {
	path := util.ReleaseDir + args.App

	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("%+v", err)
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		log.Fatalf("%+v", err)
	}

	reFileBytes := make([]byte, info.Size())
	file.Read(reFileBytes)

	return reFileBytes
}
