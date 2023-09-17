package img

import (
	"image"
	"image/color"
	"image/draw"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"github.com/nfnt/resize"
	u "github.com/zaigie/zmkx-go/internal/utils"
)

func LoadImage(filename string, threshold uint16) ([]byte, error) {
	height, width := u.EinkHeight, u.EinkWidth
	if threshold == 0 {
		threshold = 32768
	}
	// 1. 从文件中读取图片
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	// 2. 创建一个白底的画布
	canvas := image.NewRGBA(image.Rect(0, 0, width, height))
	white := color.RGBA{255, 255, 255, 255}
	draw.Draw(canvas, canvas.Bounds(), &image.Uniform{white}, image.Point{}, draw.Src)

	// 3. 缩放和绘制图片到中间位置
	bounds := img.Bounds()
	w, h := bounds.Dx(), bounds.Dy()
	var scaleFactor float64
	if w > width || h > height {
		scaleFactor = u.MinFloat(float64(width)/float64(w), float64(height)/float64(h))
	} else {
		scaleFactor = 1.0
	}
	scaledW := int(float64(w) * scaleFactor)
	scaledH := int(float64(h) * scaleFactor)
	scaledImg := resize.Resize(uint(scaledW), uint(scaledH), img, resize.Lanczos3)
	targetRect := image.Rect((width-scaledW)/2, (height-scaledH)/2, (width+scaledW)/2, (height+scaledH)/2)
	draw.Draw(canvas, targetRect, scaledImg, image.Point{}, draw.Over)

	// 4. 根据阈值，转化图片为纯黑白
	threshold32 := uint32(threshold)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r, g, b, _ := canvas.At(x, y).RGBA()
			avg := (r + g + b) / 3
			if avg < threshold32 {
				canvas.Set(x, y, color.Black)
			} else {
				canvas.Set(x, y, color.White)
			}
		}
	}

	// 5. 将处理后的图片转化为1位每像素的字节切片
	result := make([]byte, width*height/8)
	index := 0
	var currentByte byte
	var bitPosition uint8 = 0

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r, g, b, _ := canvas.At(x, y).RGBA()
			avg := (r + g + b) / 3
			if avg >= threshold32 {
				currentByte |= (1 << (7 - bitPosition))
			}

			bitPosition++
			if bitPosition == 8 {
				result[index] = currentByte
				index++
				bitPosition = 0
				currentByte = 0
			}
		}
	}

	return result, nil
}
