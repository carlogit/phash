package phash

import (
	"os"
	"testing"
)

func TestGetHash(t *testing.T) {
	file, err := os.Open("/home/carlo/Downloads/test.jpg")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	phash := GetHash(file)

	if phash != "0010011100100010000001000101001000101110100101110" {
		t.Error("expected phash is not returned: " + phash)
	}
}

func TestDistanceZero(t *testing.T) {
	hash1 := "0010011100100010000001000101001000101110100101110"
	hash2 := "0010011100100010000001000101001000101110100101110"
	
	distance := GetDistance(hash1, hash2)

	if distance != 0 {
		t.Errorf("expected distance 0, actual value: %d", distance)
	}
}

func TestDistanceTwo(t *testing.T) {
	hash1 := "0010011100100000000001000101001000101110100101110"
	hash2 := "0010011100100010000001000101001000101110100101111"
	
	distance := GetDistance(hash1, hash2)

	if distance != 2 {
		t.Errorf("expected distance 2, actual value: %d", distance)
	}
}
