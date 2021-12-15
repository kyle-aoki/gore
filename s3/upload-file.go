package s3

import (
	"bytes"
	"fmt"
	"gore/config"
	"gore/model"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"

	_ "github.com/go-sql-driver/mysql"
)

func UploadReleaseToS3(res model.Release) {
	fmt.Println("Uploading new release to S3...")
	s := GetS3Session()

	releaseFileBytes := ReadReleaseFile()
	s3FullPath := "release/" + res.App + "/" + res.QualifiedName()

	_, err := s3.New(s).PutObject(&s3.PutObjectInput{
		Bucket: aws.String(config.Gore.S3_BUCKET),
		Key:    aws.String(s3FullPath),
		Body:   bytes.NewReader(releaseFileBytes),
	})

	if err != nil {
		log.Fatalf("%+v", err)
	}
}
