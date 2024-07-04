package agent

import (
	"fmt"
	"gocv.io/x/gocv"
	"os"
	"path/filepath"
	"sync"
)

type Progress struct {
	mu       sync.Mutex
	Progress map[string]float64
}

var GlobalProgress = Progress{Progress: make(map[string]float64)}

func ProcessImages(inputDir, outputDir string, id string) {
	defer func() {
		GlobalProgress.mu.Lock()
		GlobalProgress.Progress[id] = 100
		GlobalProgress.mu.Unlock()
	}()

	img1 := gocv.IMRead(filepath.Join(inputDir, "original.jpg"), gocv.IMReadGrayScale)
	img2 := gocv.IMRead(filepath.Join(inputDir, "target.jpg"), gocv.IMReadGrayScale)
	if img1.Empty() || img2.Empty() {
		fmt.Println("Error reading images")
		return
	}

	GlobalProgress.mu.Lock()
	GlobalProgress.Progress[id] = 50
	GlobalProgress.mu.Unlock()

	psnr := CalculatePSNR(img1, img2)
	ssim := CalculateSSIM(img1, img2)

	GlobalProgress.mu.Lock()
	GlobalProgress.Progress[id] = 75
	GlobalProgress.mu.Unlock()

	result := fmt.Sprintf("PSNR: %f, SSIM: %f\n", psnr, ssim)
	err := os.WriteFile(filepath.Join(outputDir, "result.txt"), []byte(result), 0644)
	if err != nil {
		fmt.Println("Error writing result:", err)
		return
	}

	GlobalProgress.mu.Lock()
	GlobalProgress.Progress[id] = 100
	GlobalProgress.mu.Unlock()
}
