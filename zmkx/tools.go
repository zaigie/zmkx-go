package zmkx

import (
	"encoding/json"
	"image"
	"image/color"
	"image/draw"
	_ "image/jpeg"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/nfnt/resize"
)

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func proto2Map(a any) map[string]interface{} {
	jsonBytes, err := json.Marshal(a)
	if err != nil {
		log.Fatal(err)
	}
	var m map[string]interface{}
	err = json.Unmarshal(jsonBytes, &m)
	if err != nil {
		log.Fatal(err)
	}

	return m
}

func genImageID() *uint32 {
	rand.NewSource(time.Now().UnixNano())
	num := uint32(rand.Intn(900000) + 100000)
	return &num
}

func loadImage(filename string, threshold uint16) ([]byte, error) {
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

	// 2. 创建一个白底的128x296画布
	canvas := image.NewRGBA(image.Rect(0, 0, 128, 296))
	white := color.RGBA{255, 255, 255, 255}
	draw.Draw(canvas, canvas.Bounds(), &image.Uniform{white}, image.Point{}, draw.Src)

	// 3. 缩放和绘制图片到中间位置
	bounds := img.Bounds()
	w, h := bounds.Dx(), bounds.Dy()
	var scaleFactor float64
	if w > 128 || h > 296 {
		scaleFactor = min(128.0/float64(w), 296.0/float64(h))
	} else {
		scaleFactor = 1.0
	}
	scaledW := int(float64(w) * scaleFactor)
	scaledH := int(float64(h) * scaleFactor)
	scaledImg := resize.Resize(uint(scaledW), uint(scaledH), img, resize.Lanczos3)
	targetRect := image.Rect((128-scaledW)/2, (296-scaledH)/2, (128+scaledW)/2, (296+scaledH)/2)
	draw.Draw(canvas, targetRect, scaledImg, image.Point{}, draw.Over)

	// 4. 根据阈值，转化图片为纯黑白
	threshold32 := uint32(threshold)
	for y := 0; y < 296; y++ {
		for x := 0; x < 128; x++ {
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
	result := make([]byte, 128*296/8)
	index := 0
	var currentByte byte
	var bitPosition uint8 = 0

	for y := 0; y < 296; y++ {
		for x := 0; x < 128; x++ {
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
