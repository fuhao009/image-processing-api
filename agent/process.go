package agent

import (
	"fmt"
	"gocv.io/x/gocv"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

type Progress struct {
	Mu       sync.Mutex
	Progress map[string]float64
}

var GlobalProgress = Progress{Progress: make(map[string]float64)}

func ProcessImages(inputDir, outputDir string, id string) {
	defer func() {
		GlobalProgress.Mu.Lock()
		GlobalProgress.Progress[id] = 100
		GlobalProgress.Mu.Unlock()
	}()

	inputFiles, err := ioutil.ReadDir(inputDir)
	if err != nil {
		fmt.Println("Error reading input directory:", err)
		return
	}

	outputFiles, err := ioutil.ReadDir(outputDir)
	if err != nil {
		fmt.Println("Error reading output directory:", err)
		return
	}

	fileSet := make(map[string]bool)
	for _, file := range inputFiles {
		fileSet[file.Name()] = true
	}

	totalFiles := 0
	for _, file := range outputFiles {
		if _, exists := fileSet[file.Name()]; exists {
			totalFiles++
		}
	}

	if totalFiles == 0 {
		fmt.Println("No matching files to process")
		return
	}

	currentFile := 0
	for _, file := range outputFiles {
		if _, exists := fileSet[file.Name()]; exists {
			img1 := gocv.IMRead(filepath.Join(inputDir, file.Name()), gocv.IMReadGrayScale)
			img2 := gocv.IMRead(filepath.Join(outputDir, file.Name()), gocv.IMReadGrayScale)
			if img1.Empty() || img2.Empty() {
				fmt.Println("Error reading images:", file.Name())
				continue
			}

			psnr := CalculatePSNR(img1, img2)
			ssim := CalculateSSIM(img1, img2)

			result := fmt.Sprintf("Image: %s, PSNR: %f, SSIM: %f\n", file.Name(), psnr, ssim)
			err := os.WriteFile(filepath.Join(outputDir, file.Name()+".result.txt"), []byte(result), 0644)
			if err != nil {
				fmt.Println("Error writing result for", file.Name(), ":", err)
				continue
			}

			currentFile++
			progress := (float64(currentFile) / float64(totalFiles)) * 100

			GlobalProgress.Mu.Lock()
			GlobalProgress.Progress[id] = progress
			GlobalProgress.Mu.Unlock()
		}
	}
}
