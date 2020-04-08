package loop

import "fmt"

func RangeBug(){
	ch := make(chan *int, 5)

	//sender
	input := []int{1,2,3,4,5}

	go func(){
		for k, _:= range input {
			//ch <- &v
			ch <- &input[k]
		}
		close(ch)
	}()
	//receiver
	for v := range ch {
		fmt.Println(*v)
	}
}
