# Base64 Converter

This is a simple Go program that converts an image file to Base64 and saves the Base64 representation in a text file. It provides an easy way to generate a compact Base64 representation of an image.

## Prerequisites

- Go (version 1.13 or higher) installed on your machine

## Usage

1. Clone the repository or download the project files.

2. Open a terminal and navigate to the project directory.

3. Run the following command to build the Go executable:

```bash
go build
```

4. Run the executable using the following command:

```bash
go run base64-converter
```

5. You will be prompted to enter the path to the image file you want to convert. Provide the full path including the file extension.

6. If the provided file path is valid, the program will convert the image to Base64 and save the Base64 representation in a text file within the "base64" folder. You will see a success message indicating that the conversion was successful.

7. If the provided file path is invalid, an error message will be displayed, and you will be prompted to enter the path again. Repeat step 5 with a valid file path.

8. After the conversion is complete, you can find the Base64 text file in the "base64" folder within the project directory. The file will have the same name as the original image file, with the ".txt" extension.

## Notes

Only image files are supported for conversion (e.g., .png, .jpg, .jpeg, .gif).

The original image file is not modified or moved. Only the Base64 representation is saved in a separate text file.

The "base64" folder will be created automatically if it doesn't exist. Make sure you have write permissions for the project directory.

If you want to convert another image file, simply rerun the program and follow the steps again.

<strong>If you encounter any issues or have questions, feel free to open an issue in the GitHub repository.</strong>
