package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {

	args := os.Args[1:]
	count, err := strconv.Atoi(args[0])

	if err != nil {
		fmt.Println("error during conversion")
		return
	}

	for i := 0; i < count; i++ {
		go func(x int) {
			fmt.Printf("%d ", x)
		}(i)
	}

	time.Sleep(time.Second)
	fmt.Println("exiting...")
}
