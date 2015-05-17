package phash

import (
	"io"

	"github.com/disintegration/imaging"
)

const (
	size        int = 32
	smallerSize int = 8
)

func GetHash(reader io.Reader) (string, error) {
	image, err := imaging.Decode(reader)

	if err != nil {
		return "", err
	}

	image = imaging.Resize(image, size, size, imaging.Lanczos)
	image = imaging.Grayscale(image)

	dctValues := getDCTValues(image)

	return buildHash(dctValues), nil
}

func buildHash(dctValues [][]float64) string {
	dctMeanValue := calculateMeanDCTValue(dctValues)	
	
	var hash string

	for x := 0; x < smallerSize; x++ {
		for y := 0; y < smallerSize; y++ {
			if x != 0 && y != 0 {
				if dctValues[x][y] > dctMeanValue {
					hash += "1"
				} else {
					hash += "0"
				}
			}
		}
	}	
	
	return hash
}

func calculateMeanDCTValue(dctValues [][]float64) float64 {
	var total float64

	for x := 0; x < smallerSize; x++ {
		for y := 0; y < smallerSize; y++ {
			total += dctValues[x][y]
		}
	}

	total -= dctValues[0][0]

	avg := total / float64((smallerSize * smallerSize) - 1)

	return avg
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
