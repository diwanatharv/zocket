package dataaccess

import (
	"awesomeProject6/pkg/models"
	"context"
	"fmt"
	"github.com/IBM/sarama"
	"github.com/disintegration/imaging"
	"go.mongodb.org/mongo-driver/bson"
	"image"
	"image/jpeg"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func generateUniqueID() string {
	// Use the current Unix timestamp as a unique ID
	timestamp := time.Now().Unix()
	return fmt.Sprintf("%d", timestamp)
}
func updateDatabase(productID string, imagePaths []string) error {

	collection := MongoManager("product")

	filter := bson.M{"product_id": productID}

	// Define an update operation to add the imagePath to the compressed_product_images array
	update := bson.M{
		"$push": bson.M{
			"compressed_product_images": bson.M{"$each": imagePaths},
		},
	}

	// Perform the update operation

	_, err := collection.Updateone(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return nil
}
func downloadAndCompressImages(productID string) ([]string, error) {
	// Convert the productID to int64
	productIntID, err := strconv.Atoi(productID)
	fmt.Println(productIntID)
	if err != nil {
		return nil, err
	}

	collection := MongoManager("product")

	// Define a filter to find the product by ID
	filter := bson.M{"product_id": productIntID}

	// Find the product in MongoDB
	var product models.Product
	product.Updated_At = time.Now()
	err = collection.Findone(context.Background(), filter).Decode(&product)
	if err != nil {
		return nil, err
	}
	fmt.Println(product)
	var localPaths []string

	for _, imageURL := range product.ProductImages {
		// Download the image
		resp, err := http.Get(imageURL)
		if err != nil {
			log.Printf("Error downloading image: %v", err)
			continue
		}
		// Ensure the response status is OK (200)
		if resp.StatusCode != http.StatusOK {
			log.Printf("Failed to download image from %s. Status code: %d", imageURL, resp.StatusCode)
			continue
		}

		// Create a unique local file path for each image (you may want a better strategy)
		localPath := fmt.Sprintf("images/%s_%s.jpg", productID, generateUniqueID())
		localPaths = append(localPaths, localPath)

		// Create a local file to store the image
		file, err := os.Create(localPath)
		if err != nil {
			log.Printf("Error creating local file: %v", err)
			continue
		}

		// Copy the downloaded image to the local file
		_, err = io.Copy(file, resp.Body)
		if err != nil {
			log.Printf("Error copying image data: %v", err)
			continue
		}

		// Compress the image (resize and convert to JPEG)
		err = compressImage(localPath)
		if err != nil {
			log.Printf("Error compressing image: %v", err)
		}
	}

	return localPaths, nil
}
func compressImage(imagePath string) error {
	file, err := os.Open(imagePath)
	if err != nil {
		return err
	}

	// Decode the image
	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	// Resize the image (adjust dimensions as needed)
	img = imaging.Resize(img, 400, 0, imaging.Lanczos)

	// Create a new file for the compressed image
	compressedPath := strings.TrimSuffix(imagePath, ".jpg") + "_compressed.jpg"
	out, err := os.Create(compressedPath)
	if err != nil {
		return err
	}

	// Compress the image as JPEG
	err = jpeg.Encode(out, img, nil)
	if err != nil {
		return err
	}

	// Optionally, you can remove the original image if needed
	err = os.Remove(imagePath)
	if err != nil {
		log.Printf("Error removing original image: %v", err)
	}

	return nil
}
func Consumemessage(topic string, consumer sarama.Consumer) error {
	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		return err
	}

	log.Println("Kafka consumer started.")

	for {
		select {
		case msg := <-partitionConsumer.Messages():
			// Process the message here
			productID := string(msg.Value)
			log.Printf("Received product ID: %s", productID)

			// Download and compress product images
			imagePaths, err := downloadAndCompressImages(productID)
			if err != nil {
				log.Printf("Error processing product ID %s: %v", productID, err)
				return err
			}

			// Update the database with the local image paths
			err = updateDatabase(productID, imagePaths)
			if err != nil {
				log.Printf("Error updating the database for product ID %s: %v", productID, err)
				return err
			}

			log.Printf("Images processed and database updated for product ID: %s", productID)

		case err := <-partitionConsumer.Errors():
			log.Printf("Kafka consumer error: %v", err)
			return err
		}
	}
}
