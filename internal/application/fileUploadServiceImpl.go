package application

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"goFileService/internal/utils"
	"os"
	"path/filepath"
)

type FileUploadServiceImpl struct {
	// It will the aws connection
	Aws *session.Session
}

func (fileUplaod *FileUploadServiceImpl) Upload(fileName string) (string, error) {
	// The session the S3 Uploader will use

	// Create an uploader with the session and default options
	uploader := s3manager.NewUploader(fileUplaod.Aws)

	f, err := os.Open(fileName)
	if err != nil {
		return "", fmt.Errorf("failed to open file %q, %v", filepath.Base(fileName), err)
	}

	// Upload the file to S3.
	// We will use guid to generate dynamic filePath Name for the fiie to be uploaded
	fileExtension := filepath.Ext(fileName)
	fileNameForAws := fmt.Sprintf("%s%s", utils.CreateGuid(), fileExtension)
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(os.Getenv("BUCKET")),
		Key:    aws.String(fileNameForAws),
		Body:   f,
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload file, %v", err)
	}

	// We will delete the file From local
	err = os.RemoveAll(fileName)
	if err != nil {
		return "", err
	}

	return aws.StringValue(&result.Location), nil
}
