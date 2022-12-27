package teq

import (
	"math"
	"mime/multipart"
)

type FileType string

const (
	UNKNOWN FileType = ""
	JPG     FileType = "image/jpg"
	JPEG    FileType = "image/jpeg"
	PNG     FileType = "image/png"
)

func ValidateFileSize(file *multipart.FileHeader, maxSize int64) bool {
	maxSize *= int64(math.Pow(1024, 2))

	return file.Size <= maxSize
}

func ValidateImage(file *multipart.FileHeader) bool {
	return IsImage(file.Header.Get("Content-Type"))
}

func IsImage(fileType string) bool {
	switch FileType(fileType) {
	case UNKNOWN:
		return false
	case JPEG, JPG, PNG:
		return true
	default:
		return false
	}
}

func FileTypeString(fileType string) string {
	switch FileType(fileType) {
	case UNKNOWN:
		return ""
	case JPEG:
		return ".jpeg"
	case JPG:
		return ".jpg"
	case PNG:
		return ".png"
	default:
		return ""
	}
}
