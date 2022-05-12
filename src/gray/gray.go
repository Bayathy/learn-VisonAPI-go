package main

import (
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"os"
)

const threshold = 150

func main() {
	in, _ := os.Open("../../IMG_2656.jpeg")
	defer in.Close()
	out, _ := os.Create("out.jpeg")
	defer out.Close()

	Img, err := jpeg.Decode(in)
	if err != nil {
		log.Fatal(err)
	}
	srcBounds := Img.Bounds()

	// 出力用イメージ
	dest := image.NewGray(srcBounds)

	// 二値化
	for v := srcBounds.Min.Y; v < srcBounds.Max.Y; v++ {
		for h := srcBounds.Min.X; h < srcBounds.Max.X; h++ {
			c := color.GrayModel.Convert(Img.At(h, v))
			gray, _ := c.(color.Gray)
			// しきい値で二値化
			if gray.Y > threshold {
				gray.Y = 255
			} else {
				gray.Y = 0
			}
			dest.Set(h, v, gray)
		}
	}

	// 書き出し用ファイル準備
	// 書き出し
	jpeg.Encode(out, dest, nil)
}
