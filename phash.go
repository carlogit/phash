package phash

import (
	"io"

	"github.com/disintegration/imaging"
)

const (
	size        int = 32
	smallerSize int = 8
)

func GetHash(reader io.Reader) string {
	img, err := imaging.Decode(reader)

	if err != nil {
		panic(err)
	}

	img = imaging.Resize(img, size, size, imaging.Lanczos)
	img = imaging.Grayscale(img)

	dctValues := getDCTValues(img)

	return buildHash(dctValues)
}

func calculateMeanDCTValue(dctValues [][]float64) float64 {
	total := float64(0)

	for x := 0; x < smallerSize; x++ {
		for y := 0; y < smallerSize; y++ {
			total += dctValues[x][y]
		}
	}

	total -= dctValues[0][0]

	avg := total / float64((smallerSize * smallerSize) - 1)

	return avg
}

func buildHash(dctValues [][]float64) string {
	avg := calculateMeanDCTValue(dctValues)	
	
	hash := ""

	for x := 0; x < smallerSize; x++ {
		for y := 0; y < smallerSize; y++ {
			if x != 0 && y != 0 {
				if dctValues[x][y] > avg {
					hash += "1"
				} else {
					hash += "0"
				}
			}
		}
	}	
	
	return hash
}

func GetDistance(hash1, hash2 string) int {
	distance := 0
	for i := 0; i < len(hash1); i++ {
		if hash1[i] != hash2[i] {
			distance++
		}
	}

	return distance
}
