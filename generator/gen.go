package generator

import (
	"log"
	"math/rand"
	"os"
	"strconv"
)

const (
	BATCH_SIZE     = 500
	LOCATION_COUNT = 10000
)

func randomString() string {
	var bytes []byte
	for len(bytes) < 8 {
		byte := byte(rand.Intn(26) + 65)
		bytes = append(bytes, byte)
	}

	return string(bytes)
}

func randomLocations() []string {
	locsMap := make(map[string]bool, LOCATION_COUNT)
	for len(locsMap) < LOCATION_COUNT {
		new_loc := randomString()
		locsMap[new_loc] = true
	}

	var locs []string
	for loc := range locsMap {
		locs = append(locs, loc)
	}

	return locs
}

func writeFile(filepath string) *os.File {
	file, err := os.Create(filepath)
	if err != nil {
		log.Fatalf("failed to write file: %s", err)
	}

	return file
}

func streamRows(file *os.File, loc string) {
	var bytes []byte
	for i := 0; i < BATCH_SIZE; i++ {
		// append location
		bytes = append(bytes, []byte(loc)...)
		bytes = append(bytes, ';')

		// append sign
		sign := rand.Float32()
		if sign < 0.5 {
			bytes = append(bytes, '-')
		}

		// append temp
		bytes = strconv.AppendInt(bytes, rand.Int63n(100), 10)
		bytes = append(bytes, '.')
		bytes = strconv.AppendInt(bytes, rand.Int63n(10), 10)
		bytes = append(bytes, '\n')
	}
	file.Write(bytes)
}

func GenerateFile(filepath string, length int) {
	file := writeFile(filepath)
	locs := randomLocations()
	for i := 0; i < length/BATCH_SIZE; i++ {
		streamRows(file, locs[i%LOCATION_COUNT])
	}
}
