package application

type FileUploadService interface {
	// all the methods for file uploading and getting

	Upload(filePath string) error
}
