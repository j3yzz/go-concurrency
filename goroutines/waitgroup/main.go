package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Please provide a number as an argument.")
		return
	}

	count, err := strconv.Atoi(args[0])

	if err != nil {
		fmt.Printf("Error during conversion: %v\n", err)
		return
	}

	printNumbers(count)
	fmt.Println("\nExiting...")
}

func printNumbers(count int) {
	var wg sync.WaitGroup
	for i := 0; i < count; i++ {
		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			fmt.Printf("%d ", x)
		}(i)
	}

	wg.Wait()
}
