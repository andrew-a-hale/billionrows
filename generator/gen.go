package generator

import (
	"log"
	"math/rand"
	"os"
	"strconv"
)

const BATCH_SIZE = 500

var STATIONS = [10]string{
	"Brisbane",
	"Sydney",
	"Melbourne",
	"Albury",
	"Morayfield",
	"Armadale",
	"Dalby",
	"Gold Coast",
	"Aspley",
	"Casledine",
}

func writeFile(filepath string) *os.File {
	file, err := os.Create(filepath)
	if err != nil {
		log.Fatalf("failed to write file: %s", err)
	}

	return file
}

func streamRows(file *os.File) {
	var bytes []byte
	for i := 0; i < BATCH_SIZE; i++ {
		bytes = append(bytes, []byte(STATIONS[i%10])...)
		bytes = append(bytes, ';')
		bytes = strconv.AppendInt(bytes, rand.Int63n(100), 10)
		bytes = append(bytes, '\n')
	}
	file.Write(bytes)
}

func GenerateFile(filepath string, length int) {
	file := writeFile(filepath)
	for i := 0; i < length/BATCH_SIZE; i++ {
		streamRows(file)
	}
}
