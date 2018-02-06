package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

const (
	backgroundColorIndex = 0 // パレットの最初の色
	lineColorIndex = 1 // パレットの次の色
)

var black = color.RGBA{0x00, 0x00, 0x00, 0xff}
var red = color.RGBA{0xff, 0x00, 0x00, 0xff}
var green = color.RGBA{0x00, 0xff, 0x00, 0xff}
var blue = color.RGBA{0x00, 0x00, 0xff, 0xff}
var palette = []color.Color{black, red, green, blue}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles = 5 // 発振器 x が完了する周回の回数
		res    = 0.001 // 回転の分解能
		size = 100 // 画像キャンパスは [-size..+size] の範囲を扱う
		nframes = 64 // アニメーションフレーム数
		delay = 8 // 10ms 単位でのフレーム間の遅延
	)
	freq := rand.Float64() * 3.0 // 発振器 y の相対周波数
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // 位相差
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			lineColorPaletteIndex := uint8((i % 3) + 1)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), lineColorPaletteIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // 注意: エンコードエラーを無視
}
