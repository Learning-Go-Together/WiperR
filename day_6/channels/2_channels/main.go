package channels

import "fmt"

func Channels() {
	ch := make(chan string)
	go func() {
		ch <- "Hello Shivang!"
		ch <- "Hello Shivang!"
	}()
	v := <-ch
	fmt.Println(v)
}
