package main

import (
	"log"
	"os"
	"strconv"

	"github.com/andrew-a-hale/billionrows/generator"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatalln("Usage: billionrows <filepath> <length>")
	}

	filepath := os.Args[1]
	length, err := strconv.ParseInt(os.Args[2], 10, 0)
	if err != nil {
		log.Fatalln("Length was unable to be parsed")
	}

	generator.GenerateFile(filepath, int(length))
}
