package main

import (
	"fmt"
	"sync"
)

func makeThumbnails6(filenames <-chan string) int64 {
	sizes := make(chan int64)
	var wg sync.WaitGroup // number of working goroutines
	fmt.Println(1)
	for f := range filenames {
		fmt.Println(2)
		wg.Add(1)
		// worker
		go func(f string) {
			defer wg.Done()
			size := int64(len(f))
			//thumb, err := thumbnail.ImageFile(f)
			//if err != nil {
			//	log.Println(err)
			//	return
			//}
			//info, _ := os.Stat(thumb) // OK to ignore error
			fmt.Println(3)
			sizes <- size
			fmt.Println(4)
		}(f)
	}
	fmt.Println(5)
	// closer
	go func() {
		wg.Wait()
		fmt.Println(6)
		close(sizes)
	}()

	var total int64
	for size := range sizes {
		fmt.Println(7)
		total += size
	}
	fmt.Println(8)

	return total
}

func main() {
	ch := make(chan string)
	go func() {
		ch <- "interface"
		ch <- "3435"
		close(ch)
	}()

	total := makeThumbnails6(ch)
	fmt.Println("total--", total)
}
