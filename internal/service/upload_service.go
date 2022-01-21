package service

import (
	"bytes"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/uuid"
	"log"
	"os"
	"path/filepath"
)

type UploadService struct {
	s3Client *s3.S3
	bucket   string
}

func CreateUploadService(s3Client *s3.S3, bucket string) *UploadService {
	return &UploadService{
		s3Client,
		bucket,
	}
}

func CreateDefaultUploadService() *UploadService {
	s, err := session.NewSession(&aws.Config{Region: aws.String(os.Getenv("S3_REGION"))})
	if err != nil {
		log.Fatal(err)
	}
	return CreateUploadService(s3.New(s), os.Getenv("S3_BUCKET"))
}

func (u *UploadService) UploadImage(file *os.File) (string, error) {
	log.Print("sanity UploadImage file name :: ", file.Name())
	//log.Print("sanity UploadImage file ext :: ", filepath.Ext(file.()))
	fileInfo, err := file.Stat()
	if err != nil {
		log.Print("error file stat :: ", err)
		return "", err
	}
	size := fileInfo.Size()
	buffer := make([]byte, size)
	_, _ = file.Read(buffer)
	s3Key := uuid.New().String() + filepath.Ext(file.Name())
	_, err = u.s3Client.PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(u.bucket),
		Key:                  aws.String(s3Key),
		ACL:                  aws.String("public-read"),
		Body:                 bytes.NewReader(buffer),
		ContentLength:        aws.Int64(size),
		ContentType:          aws.String(getContentType(file.Name())),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
	})
	if err != nil {
		log.Print(err)
	}
	return s3Key, nil
}

func getContentType(file string) string {
	ext := filepath.Ext(file)
	if ext == ".png" {
		return "image/png"
	} else if ext == ".jpg" || ext == ".jpeg" {
		return "image/jpeg"
	} else {
		return "application/octet-stream"
	}
}
