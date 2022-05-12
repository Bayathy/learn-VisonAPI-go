package gray

import (
    "image"
    "image/png"
    "image/color"
    "os"
)

// 二値化のしきい値
const threshold = 128

func main() {
    // 画像ファイルを開く(書き込み元)
    src, _ := os.Open("../IMG_2656.jpeg")
    defer src.Close()

    // デコードしてイメージオブジェクトを準備
    srcImg, _, err := image.Decode(src)
    if err != nil {
        panic(err)
    }
    srcBounds := srcImg.Bounds()

    // 出力用イメージ
    dest := image.NewGray(srcBounds)

    // 二値化
    for v := srcBounds.Min.Y; v < srcBounds.Max.Y; v++ {
        for h := srcBounds.Min.X; h < srcBounds.Max.X; h++ {
            c := color.GrayModel.Convert(srcImg.At(h, v))
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
    outfile, _ := os.Create("out.png")
    defer outfile.Close()
    // 書き出し
    png.Encode(outfile, dest)
}