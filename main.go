package main

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// Prompt the user to enter the path to the image file
	var imagePath string
	for {
		fmt.Print("Enter the path to the image file: ")
		_, err := fmt.Scanln(&imagePath)
		if err != nil {
			log.Fatal(err)
		}

		// Check if the file exists
		_, err = os.Stat(imagePath)
		if err == nil {
			break // File exists, proceed with conversion
		}

		fmt.Println("Invalid file path. Please try again.")
	}

	// Read the image file
	imageData, err := ioutil.ReadFile(imagePath)
	if err != nil {
		log.Fatal(err)
	}

	// Convert the image to Base64
	base64Data := base64.StdEncoding.EncodeToString(imageData)

	// Create the "base64" folder if it doesn't exist
	err = os.MkdirAll("base64", 0755)
	if err != nil {
		log.Fatal(err)
	}

	// Extract the original image name and extension
	imageName := filepath.Base(imagePath)
	imageExt := filepath.Ext(imagePath)

	// Generate a unique file name using the SHA-1 hash of the image data
	filename := generateUniqueFilename(imageData)

	// Construct the new file name with the original name and hash
	newFilename := strings.TrimSuffix(imageName, imageExt) + "_" + filename + imageExt

	// Write the Base64 representation to a text file in the "base64" folder with the new file name
	err = ioutil.WriteFile(filepath.Join("base64", strings.TrimSuffix(newFilename, imageExt)+".txt"), []byte(base64Data), 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Base64 conversion successful!")
}

// Function to generate a unique file name using the SHA-1 hash of the data
func generateUniqueFilename(data []byte) string {
	hash := sha1.Sum(data)
	return fmt.Sprintf("%x", hash)
}
