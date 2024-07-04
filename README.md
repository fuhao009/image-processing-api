
# Image Processing API

This project provides an API built with Gin to process images and calculate PSNR and SSIM between two images. The API supports uploading image directories and querying the processing progress.

## Features

- Calculate the Peak Signal-to-Noise Ratio (PSNR) between two images.
- Calculate the Structural Similarity Index (SSIM) between two images.
- Track and query the progress of the image processing.

## Endpoints

### 1. Upload Image Directory and Start Processing

**Endpoint:** `/upload`  
**Method:** `POST`  
**Description:** Upload the input and output directories and start the image processing.

**Request:**

- `input_dir` (form-data): Path to the input directory containing the images.
- `output_dir` (form-data): Path to the output directory where the result will be saved.
- `id` (form-data): Unique ID for the process.

**Example cURL Request:**

```sh
curl -X POST http://localhost:8080/upload \
    -F "input_dir=/path/to/input/dir" \
    -F "output_dir=/path/to/output/dir" \
    -F "id=unique-process-id"
```

**Response:**

```json
{
  "message": "Processing started",
  "id": "unique-process-id"
}
```

### 2. Get Processing Progress

**Endpoint:** `/progress`  
**Method:** `GET`  
**Description:** Get the current progress of the image processing.

**Request Parameters:**

- `id` (query): The unique ID of the process.

**Example cURL Request:**

```sh
curl -X GET "http://localhost:8080/progress?id=unique-process-id"
```

**Response:**

```json
{
  "id": "unique-process-id",
  "progress": 75
}
```

## Setup and Run

### Prerequisites

- Go 1.22 or higher
- Gin
- GoCV

### Installation

1. Clone the repository:

```sh
git clone https://github.com/yourusername/image-processing-api.git
cd image-processing-api
```

2. Install the required dependencies:

```sh
go mod tidy
```

3. Run the API:

```sh
go run main.go
```