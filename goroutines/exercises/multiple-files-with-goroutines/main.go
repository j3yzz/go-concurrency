package main

import (
	"log"
	"os"
	"path/filepath"
	"strconv"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	start := 0
	end := 10

	basePath := "/tmp"

	for i := start; i <= end; i++ {
		wg.Add(1)
		filename := "master" + strconv.Itoa(i)
		path := filepath.Join(basePath, filename)

		go func(fp string) {
			defer wg.Done()
			createFile(fp)
		}(path)
	}

	wg.Wait()
}

func createFile(fp string) {
	file, err := os.Create(fp)
	if err != nil {
		log.Fatal(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
}
