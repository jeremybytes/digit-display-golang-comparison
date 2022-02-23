package display

import "strings"

func GetImagesAsString(image1 []int, image2 []int) string {
	firstImage := GetImageAsArray(image1)
	secondImage := GetImageAsArray(image2)

	var output strings.Builder

	for i := 0; i < 28; i++ {
		output.WriteString(firstImage[i])
		output.WriteString(" | ")
		output.WriteString(secondImage[i])
		output.WriteString("\n")
	}
	return output.String()
}

func GetImageAsArray(imageData []int) [28]string {
	var output [28]string
	var line strings.Builder
	for i, pixel := range imageData {
		if i%28 == 0 && i != 0 {
			output[(i/28)-1] = line.String()
			line.Reset()
		}
		outputChar := getDisplayCharForPixel(pixel)
		line.WriteString(outputChar)
		line.WriteString(outputChar)
	}
	output[27] = line.String()
	return output
}

func getDisplayCharForPixel(i int) string {
	switch {
	case (i > 16 && i < 32):
		return "."
	case (i >= 32 && i < 64):
		return ":"
	case (i >= 64 && i < 160):
		return "o"
	case (i >= 160 && i < 224):
		return "O"
	case (i >= 224):
		return "@"
	default:
		return " "
	}
}
