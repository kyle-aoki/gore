package s3

import (
	"gore/config"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

func GetS3Session() *session.Session {
	s, err := session.NewSession(&aws.Config{Region: aws.String(config.Gore.S3_REGION)})
	if err != nil {
		log.Fatal(err)
	}
	return s
}
