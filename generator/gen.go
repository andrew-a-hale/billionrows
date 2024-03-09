package generator

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"runtime"
	"strconv"
	"sync"
)

const LOCATION_COUNT = 10000

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

func generateData(locs []string, lines int) []byte {
	var bytes []byte
	for i := 0; i < lines; i++ {
		// append location
		loc := locs[i%LOCATION_COUNT]
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

	return bytes
}

func GenerateFile(filepath string, length int) error {
	file := writeFile(filepath)
	defer file.Close()
	locs := randomLocations()
	maxGoRoutines := runtime.GOMAXPROCS(0)
	chunkSize := length / 100 / maxGoRoutines
	var wg sync.WaitGroup

	written := 0
	for written < length {
		for i := 0; i < maxGoRoutines; i++ {
			wg.Add(1)
			startLine := written
			endLine := min(startLine+chunkSize, length)

			go func(startLine int, endLine int, locs []string, file *os.File, wg *sync.WaitGroup) {
				writer := bufio.NewWriter(file)
				writer.Write(generateData(locs, endLine-startLine))
				writer.Flush()
				wg.Done()
			}(startLine, endLine, locs, file, &wg)

			written += chunkSize
		}
		wg.Wait()
	}

	return nil
}
