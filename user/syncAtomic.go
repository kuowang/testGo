package main

/*
import (
	"fmt"
	"sync/atomic"
	"time"
)

var i int32 = 100

func add() {
	atomic.AddInt32(&i, 1)
}
func sub() {
	atomic.AddInt32(&i, -1)
}
func main() {

	for i := 0; i < 100; i++ {
		go add()
		go sub()
	}
	time.Sleep(time.Second * 2)
	fmt.Println("end ...", i)
	atomic.LoadInt32(&i)       //读的时候查看值锁
	atomic.StoreInt32(&i, 200) //写的时候查看原子操作
	atomic.CompareAndSwapInt32(&i, 100, 200)

}
*/
