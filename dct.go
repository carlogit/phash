package phash

import (
	"image"
	"math"
)

var c []float64 = initCoefficients()

func initCoefficients() []float64 {
	c := make([]float64, size)

	for i := 1; i < size; i++ {
		c[i] = 1
	}

	c[0] = 1 / math.Sqrt(2.0)
	
	return c
}

func getInitDCTTable(img image.Image) [][]float64 {
	xSize := img.Bounds().Max.X
	ySize := img.Bounds().Max.Y

	vals := make([][]float64, xSize)

	for x := 0; x < xSize; x++ {
		vals[x] = make([]float64, ySize)
		for y := 0; y < ySize; y++ {
			vals[x][y] = getBlue(img, x, y)
		}
	}

	return vals
}

func getBlue(img image.Image, x int, y int) float64 {
	r, _, _, _ := img.At(x, y).RGBA()
	return float64(r & 0xff)
}

func getDCTValues(img image.Image) [][]float64 {
	vals := getInitDCTTable(img)
	N := size

	F := make([][]float64, N)
	for u := 0; u < N; u++ {
		F[u] = make([]float64, N)
		for v := 0; v < N; v++ {
			sum := float64(0.0)
			for i := 0; i < N; i++ {
				for j := 0; j < N; j++ {
					valor := math.Cos((2.0*float64(i)+1.0)/(2.0*float64(N))*float64(u)*math.Pi) * math.Cos((2.0*float64(j)+1.0)/(2.0*float64(N))*float64(v)*math.Pi) * (vals[i][j])
					sum += valor
				}
			}
			sum *= ((c[u] * c[v]) / 4.0)
			F[u][v] = sum
		}
	}
	return F
}
