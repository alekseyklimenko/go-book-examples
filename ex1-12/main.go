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
	"strconv"
	"time"
)

var palette = []color.Color{
	color.Black,
	color.RGBA{0, 0xff, 0, 0xff},
	color.RGBA{0xff, 0, 0, 0xff},
	color.RGBA{0, 0, 0xff, 0xff},
	color.RGBA{0xff, 0xff, 0, 0xff},
	color.RGBA{0, 0xff, 0xff, 0xff},
	color.RGBA{0xff, 0xff, 0xff, 0xff},
}

const maxColorIndex = 6

type imageParams struct {
	cycles  float64 // number of complete x oscillator revolutions
	res     float64 // angular resolution
	size    int     // image canvas covers [-size..+size]
	nframes int     // number of animation frames
	delay   int     // delay between frames in 10ms units
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the HTTP request.
func handler(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UTC().UnixNano())
	params := imageParams{5, 0.001, 100, 64, 8}
	if err := r.ParseForm(); err == nil {
		setupParams(r, &params)
	}
	lissajous(w, params)
}

func setupParams(r *http.Request, params *imageParams) {
	for k, v := range r.Form {
		value, err := strconv.Atoi(v[0])
		if err != nil {
			continue
		}
		switch k {
		case "cycles":
			params.cycles = float64(value)
		case "res":
			params.res = float64(value)
		case "size":
			params.size = value
		case "nframes":
			params.nframes = value
		case "delay":
			params.delay = value
		}
	}
}

func lissajous(out io.Writer, params imageParams) {
	var colorIndex uint8
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: params.nframes}
	phase := 0.0 // phase difference
	size := params.size
	sizef := float64(size)
	for i := 0; i < params.nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < params.cycles*2*math.Pi; t += params.res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			colorIndex = uint8(i%maxColorIndex + 1)
			img.SetColorIndex(size+int(x*sizef+0.5), size+int(y*sizef+0.5), colorIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, params.delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
