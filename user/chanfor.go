package main

/*
import "fmt"

var c = make(chan int)

func main() {

	go func() {
		for i := 0; i < 20; i++ {
			c <- i
		}
		close(c) //如果不关闭  则读的次数超出通道数量，则出现死锁
	}()

	//循环方式1
	r := <-c
	fmt.Println(r)
	r = <-c
	fmt.Println(r)

	//循环方式2
	for i := 0; i < 2; i++ {
		r := <-c
		fmt.Println(r)
	}
	//循环方式3
	for v := range c {
		fmt.Println(v)
	}
	//循环方式4
	for {
		v, ok := <-c
		if ok {
			fmt.Println(v)
		} else {
			break
		}
	}

}
*/
