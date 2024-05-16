package utils

import (
	"log"
	"mime/multipart"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/joho/godotenv"
)

var s3Client *s3.S3

func InitS3() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	s3Config := &aws.Config{
		Credentials: credentials.NewStaticCredentials(
			os.Getenv("S3_ACCESS_KEY"),
			os.Getenv("S3_SECRET_KEY"),
			""),
		Endpoint: aws.String(os.Getenv("S3_ENDPOINT")),
		Region:   aws.String("sgp1"),
	}

	sess, err := session.NewSession(s3Config)
	if err != nil {
		log.Fatalf("Failed to initialize new session: %v", err)
	}

	s3Client = s3.New(sess)
}

func UploadFileToS3(file *multipart.FileHeader, folderName string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	key := folderName + "/" + file.Filename

	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String("go-chat"),
		Key:    aws.String(key),
		Body:   src,
		ACL:    aws.String("public-read"),
	})

	if err != nil {
		return err
	}

	return nil
}
