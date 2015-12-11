package main

import (
	"os"
	"image"
	"image/jpeg"
	"image/draw"
	"fmt"
	"image/color"
)

func Coloring(img *image.NRGBA, x0 int, x1 int, y0 int, y1 int, set_color color.NRGBA) {
	for  x:= x0; x < x1; x++ {
		for y := y0; y < y1; y ++ {
			img.Set(x, y, set_color)
		}
	}
}

func AddPhtoFrame() {
	fimage, _ := os.Open("res/gophergala.jpg")
	defer fimage.Close()
	img, _, _ := image.Decode(fimage)
	w, h := img.Bounds().Max.X, img.Bounds().Max.Y
	img2 := image.NewNRGBA(img.Bounds())
	draw.Draw(img2, img2.Bounds(), img, image.Point{0,0}, draw.Src)
	Coloring(img2, 0, w, 0, 100, color.NRGBA{255, 255, 0, 255})
	Coloring(img2, 0, w, 100, 120, color.NRGBA{255, 0, 0, 255})
	Coloring(img2, 0, w, h-120, h-100, color.NRGBA{255, 255, 0, 255})
	Coloring(img2, 0, w, h-100, h, color.NRGBA{255, 0, 0, 255})
	Coloring(img2, 0, 80, 0, h, color.NRGBA{0, 0, 255, 255})
	Coloring(img2, 80, 100, 0, h, color.NRGBA{0, 255, 0, 255})
	Coloring(img2, w-100, w-80, 0, h, color.NRGBA{0, 0, 255, 255})
	Coloring(img2, w-80, w, 0, h, color.NRGBA{0, 255, 0, 255})
	toimg, _ := os.Create("new.jpg")
	defer toimg.Close()
	jpeg.Encode(toimg, img2, &jpeg.Options{jpeg.DefaultQuality})
}

func main() {
	AddPhtoFrame()
	println(`This final exercise,let's add a photo frame for gala logo!
You should use image package to generate a new iamge from the gala log(which is stored in res/gophergala.jpg,and makes it like the res/m.jpg.
Now edit main.go to complete 'AddPhtoFrame' function,this task has no test,enjoy your trip!`)
}
