package dataaccess

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateUniqueID(t *testing.T) {
	uniqueID := generateUniqueID()

	// Check if the generated ID is not empty
	assert.NotEmpty(t, uniqueID)
}
func TestCompressImage(t *testing.T) {
	// Assuming you have an image file named "test.jpg" in the current directory
	imagePath := "test.jpg"

	err := compressImage(imagePath)

	// Check if there was no error during compression
	assert.NoError(t, err)
}
func TestDownloadAndCompressImages(t *testing.T) {
	// Assuming you have a valid product ID for testing
	productID := "1"

	imagePaths, err := downloadAndCompressImages(productID)

	// Check if there was no error during image download and compression
	assert.NoError(t, err)

	// Check if the imagePaths slice is not empty
	assert.NotEmpty(t, imagePaths)
}
