package loop

import "fmt"

// RangeBug range bugs
func RangeBug() {
	ch := make(chan *int, 5)

	//sender
	input := []int{1, 2, 3, 4, 5}

	go func(input []int) {
		for _, v := range input {
			ch <- &v
			//ch <- &input[k]
		}
		close(ch)
	}(input)
	//receiver
	for v := range ch {
		fmt.Println(*v)
	}
}
