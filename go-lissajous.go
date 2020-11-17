package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"time"
)

var blackColor = color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xFF}
var yellowColor = color.RGBA{R: 0xFF, G: 0xEB, B: 0x00, A: 0xff}
var vividRedColor = color.RGBA{R: 0xFC, G: 0x00, B: 0x19, A: 0xff}
var malachiteColor = color.RGBA{R: 0x01, G: 0xFF, B: 0x4F, A: 0xff}
var shockingPinkColor = color.RGBA{R: 0xFF, G: 0x01, B: 0xD7, A: 0xff}
var interdimensionalBlueColor = color.RGBA{R: 0x56, G: 0x00, B: 0xCC, A: 0xff}
var turquoiseBlueColor = color.RGBA{R: 0x00, G: 0xED, B: 0xF5, A: 0xff}

var palette = color.Palette{
	yellowColor,
	vividRedColor,
	malachiteColor,
	shockingPinkColor,
	interdimensionalBlueColor,
	turquoiseBlueColor}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	lissajous(w)
}

func lissajous(out io.Writer) {
	shufflePalette()

	anim := generateAnimation()

	gif.EncodeAll(out, &anim)
}

func shufflePalette() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(palette), func(i, j int) { palette[i], palette[j] = palette[j], palette[i] })
}

func generateAnimation() gif.GIF {
	const (
		cycles = 5
		res = 0.001
		size = 200
		nFrames = 64
		delay = 16
	)

	freq := rand.Float64() * 5.0

	anim := gif.GIF{LoopCount: nFrames}
	colorIndex := uint8(rand.Intn(3) + 1)
	phase := 0.0

	for i := 0; i < nFrames; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			setImgColorIndex(img, x, y, colorIndex)
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	return anim
}

func setImgColorIndex(img *image.Paletted, x float64, y float64, colorIndex uint8) {
	const size = 200
	xIndex := size+int(x*size)
	yIndex := size+int(y*size)

	// Set color index on 9x9 block
	for i := -1; i < 1; i++ {
		for j := -1; j < 1; j++ {
			img.SetColorIndex(xIndex+i, yIndex+j, colorIndex)
		}
	}
}