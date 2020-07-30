package handler

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/caarlos0/env"
	"io"
)

type EnvConfig struct {
	AccessKey         string        `env:"ACCESS_KEY"`
	AccessSecret      string		`env:"ACCESS_SECRET"`
	DbLocation        string		`env:"DB_LOCATION"`
	Region            string		`env:"REGION"`
	Bucket            string		`env:"BUCKET"`
	Port			  string 		`env:"PORT"`
}

var AwsUploader *s3manager.Uploader
var Envs = EnvConfig{}
var awsConfig *aws.Config

func AwsConfig() {
	if err := env.Parse(&Envs); err != nil {
		fmt.Printf("%+v\n", err)
	}

	awsConfig = &aws.Config{
		Region:      aws.String(Envs.Region),
		Credentials: credentials.NewStaticCredentials(Envs.AccessKey, Envs.AccessSecret, ""),
	}

	sess := session.Must(session.NewSession(awsConfig))
	uploader := s3manager.NewUploader(sess, func(u *s3manager.Uploader) {
		u.PartSize = 10 * 1024 * 1024
		u.Concurrency = 2
	})

	AwsUploader = uploader
}

func FileUploadTos3(key string, file io.Reader) (*s3manager.UploadOutput, error) {
	if err := env.Parse(&Envs); err != nil {
		fmt.Printf("%+v\n", err)
	}

	awsResponse, awsErr := AwsUploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(Envs.Bucket),
		Key:    aws.String("/photo/" + key),
		Body:   file,
	})

	return awsResponse, awsErr
}
