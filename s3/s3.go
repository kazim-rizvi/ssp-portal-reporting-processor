// s3/s3.go

package s3

import (
	"fmt"
	"ssp-portal-reporting-processor/config"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3Uploader struct {
	cfg config.S3Config
}

func NewS3Uploader(cfg config.S3Config) *S3Uploader {
	return &S3Uploader{cfg: cfg}
}

func (su *S3Uploader) UploadFileAndGeneratePresignedURL(filePath string) (string, error) {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(su.cfg.Region),
		Credentials: credentials.NewStaticCredentials(su.cfg.AccessKey, su.cfg.SecretKey, ""),
	})
	if err != nil {
		return "", err
	}

	s3Client := s3.New(sess)

	// Upload the file to S3
	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(su.cfg.Bucket),
		Key:    aws.String("report.csv"), // Adjust the object key as needed
		Body:   aws.ReadSeekCloser(strings.NewReader(filePath)),
	})
	if err != nil {
		return "", err
	}

	// Calculate the expiry time for the presigned URL using the configured duration
	expiryDuration := time.Duration(su.cfg.PresignedURLExpiryHours) * time.Hour

	req, _ := s3Client.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(su.cfg.Bucket),
		Key:    aws.String("report.csv"), // Adjust the object key as needed
	})
	presignedURL, err := req.Presign(expiryDuration)
	if err != nil {
		return "", err
	}

	fmt.Println("File uploaded to S3, and presigned URL generated successfully!")
	return presignedURL, nil
}
