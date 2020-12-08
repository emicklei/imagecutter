package main

import (
	"flag"
	"image/png"
	"log"
	"os"

	"github.com/oliamb/cutter"
)

var (
	oWidth  = flag.Int("w", 950, "width in pixels")
	oHeight = flag.Int("h", 600, "height in pixels")
	oInput  = flag.String("i", "", "PNG file")
	oOutput = flag.String("o", "", "PNG file, overwrites source if empty")
)

func main() {
	flag.Parse()
	if len(*oInput) == 0 {
		flag.Usage()
		return
	}
	in, err := os.Open(*oInput)
	check(err)
	img, err := png.Decode(in)
	check(err)
	in.Close()
	croppedImg, err := cutter.Crop(img, cutter.Config{
		Width:  *oWidth,
		Height: *oHeight,
		Mode:   cutter.Centered,
	})
	check(err)
	pngOut := *oOutput
	if len(pngOut) == 0 {
		pngOut = *oInput
	}
	out, err := os.Create(pngOut)
	check(err)
	err = png.Encode(out, croppedImg)
	check(err)
	out.Close()
}

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
