package application

import "github.com/aws/aws-sdk-go/aws/session"

type FileUploadServiceImpl struct {
	// It will the aws connection
	Aws *session.Session
}

func (fileUplaod *FileUploadServiceImpl) Upload(filePath string) error {
	return nil
}
