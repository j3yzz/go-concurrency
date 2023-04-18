package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	fmt.Printf("%#v\n", wg)

	args := os.Args[1:]
	count, err := strconv.Atoi(args[0])

	if err != nil {
		fmt.Println("error during conversion")
		return
	}

	for i := 0; i < count; i++ {
		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			fmt.Printf("%d ", x)
		}(i)
	}
	fmt.Printf("\n")
	fmt.Printf("%#v\n", wg)
	wg.Wait()
	fmt.Println("exiting...")
}
