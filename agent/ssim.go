package agent

//
//import (
//	"fmt"
//	"gocv.io/x/gocv"
//	"math"
//)
//
//// Mean calculates the mean of an image
//func Mean(img gocv.Mat) float64 {
//	sum := 0.0
//	for y := 0; y < img.Rows(); y++ {
//		for x := 0; x < img.Cols(); x++ {
//			sum += float64(img.GetUCharAt(y, x))
//		}
//	}
//	return sum / float64(img.Total())
//}
//
//// Variance calculates the variance of an image
//func Variance(img gocv.Mat, mean float64) float64 {
//	var sum float64
//	for y := 0; y < img.Rows(); y++ {
//		for x := 0; x < img.Cols(); x++ {
//			value := float64(img.GetUCharAt(y, x))
//			sum += math.Pow(value-mean, 2)
//		}
//	}
//	return sum / float64(img.Total())
//}
//
//// Covariance calculates the covariance between two images
//func Covariance(img1, img2 gocv.Mat, mean1, mean2 float64) float64 {
//	var sum float64
//	for y := 0; y < img1.Rows(); y++ {
//		for x := 0; x < img1.Cols(); x++ {
//			value1 := float64(img1.GetUCharAt(y, x))
//			value2 := float64(img2.GetUCharAt(y, x))
//			sum += (value1 - mean1) * (value2 - mean2)
//		}
//	}
//	return sum / float64(img1.Total())
//}
//
//// CalculateSSIM calculates the Structural Similarity Index between two images
//func CalculateSSIM(img1, img2 gocv.Mat) float64 {
//	mean1 := Mean(img1)
//	mean2 := Mean(img2)
//	var1 := Variance(img1, mean1)
//	var2 := Variance(img2, mean2)
//	cov := Covariance(img1, img2, mean1, mean2)
//
//	C1 := math.Pow(0.01*255, 2)
//	C2 := math.Pow(0.03*255, 2)
//
//	ssim := ((2*mean1*mean2 + C1) * (2*cov + C2)) / ((math.Pow(mean1, 2) + math.Pow(mean2, 2) + C1) * (var1 + var2 + C2))
//	return ssim
//}
//
//// CalculatePSNR calculates the Peak Signal-to-Noise Ratio between two images
//func CalculatePSNR(img1, img2 gocv.Mat) float64 {
//	if img1.Rows() != img2.Rows() || img1.Cols() != img2.Cols() {
//		return 0.0
//	}
//
//	var mse float64
//	for y := 0; y < img1.Rows(); y++ {
//		for x := 0; x < img1.Cols(); x++ {
//			diff := img1.GetUCharAt(y, x) - img2.GetUCharAt(y, x)
//			mse += float64(diff * diff)
//		}
//	}
//
//	mse /= float64(img1.Rows() * img1.Cols())
//	if mse == 0 {
//		return math.Inf(1)
//	}
//
//	return 20 * math.Log10(255.0/math.Sqrt(mse))
//}
//
//func Ssim() {
//	img1 := gocv.IMRead("assets/原图.jpg", gocv.IMReadGrayScale)
//	img2 := gocv.IMRead("assets/恒生.jpg", gocv.IMReadGrayScale)
//	if img1.Empty() || img2.Empty() {
//		fmt.Println("Error reading images")
//		return
//	}
//
//	psnr := CalculatePSNR(img1, img2)
//	ssim := CalculateSSIM(img1, img2)
//
//	fmt.Printf("PSNR: %f, SSIM: %f\n", psnr, ssim)
//}
