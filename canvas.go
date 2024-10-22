package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"os/exec"
)

var myColors = struct {
	black color.NRGBA
	white color.NRGBA
	red   color.NRGBA
	green color.NRGBA
	blue  color.NRGBA
}{
	black: color.NRGBA{R: 0, G: 0, B: 0, A: 255},
	white: color.NRGBA{R: 255, G: 255, B: 255, A: 255},
	red:   color.NRGBA{R: 255, G: 0, B: 0, A: 255},
	green: color.NRGBA{R: 0, G: 255, B: 0, A: 255},
	blue:  color.NRGBA{R: 0, G: 0, B: 255, A: 255},
}

func CanvasInit(size int, color1 color.NRGBA) *image.NRGBA {
	rect := image.Rect(0, 0, size, size)
	img := image.NewNRGBA(rect)

	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			img.SetNRGBA(x, y, color1)
		}
	}

	return img
}

func SaveImage(img *image.NRGBA) {
	// Encode as PNG and write to stdout
	f, err := os.Create("image.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if err := png.Encode(f, img); err != nil {
		panic(err)
	}

	// Open the image using the default image viewer
	cmd := exec.Command("open", "image.png")
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}

func IncreaseLineThickness(canvas *image.NRGBA, x, y, thickness int, color color.NRGBA) {
	size := thickness / 2

	for i := 0; i < size; i++ {
		for n := 0; n < size; n++ {
			canvas.SetNRGBA(x+i+n, y+i, color)
			canvas.SetNRGBA(x-i+n, y+i, color)
		}
	}
}

func Draw(canvas *image.NRGBA, points []Point, thickness int, color color.NRGBA) {
	for _, v := range points {
		x := v.x
		y := v.y
		IncreaseLineThickness(canvas, int(x), int(y), thickness, color)
	}
}

func DrawAxis(img *image.NRGBA, rotation float32, thickness int) {

	size := img.Bounds().Max.X
	bias := size / 2

	pX := Polynomy{coefficient: []float32{rotation - 1}, bias: float32(bias)}

	pY := Polynomy{coefficient: []float32{0}, bias: float32(bias)}

	// always start in isometric view
	pZ := Polynomy{coefficient: []float32{rotation}, bias: float32(bias)}

	x := CreateArray(size, 0) //creates an array of size 'size'
	x1 := Map(x, pX.f_p, size/2)
	y := Map(x, pY.f_p, 0)
	z := Map(x, pZ.f_p, size/2)

	pointsX := make([]Point, len(x))
	pointsY := make([]Point, len(x))
	pointsZ := make([]Point, len(x))

	for i := range x {
		pointsX[i] = Point{x: x[i], y: x1[i]}
		pointsY[i] = Point{x: y[i], y: x[i]}
		pointsZ[i] = Point{x: x[i], y: z[i]}
	}

	Draw(img, pointsX, thickness, myColors.green)
	Draw(img, pointsY, thickness, myColors.red)
	Draw(img, pointsZ, thickness, myColors.blue)
	// Cartesian map printed

}

func main() {
	size := 2000
	thickness := 10
	bias := float32(size / 2)

	img := CanvasInit(size, myColors.white)
	rotation := float32(0.4)
	DrawAxis(img, rotation, thickness)

	// Custom function
	p := Polynomy{
		coefficient: []float32{rotation - 1, -0.01},
		bias:        float32(bias),
	}

	x := CreateArray(size, 0) //creates an array of size 'size'

	biasPX := size / 2
	pointsP := make([]Point, len(x))
	pF := Map(x, p.f_p, biasPX)

	for i := range x {
		pointsP[i] = Point{x: x[i], y: pF[i]}
	}

	Draw(img, pointsP, thickness, myColors.black)

	SaveImage(img)
}
