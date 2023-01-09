package main

/*
import (
	"fmt"
	"time"
)

func main() {

	timer := time.NewTimer(time.Second * 2) //延迟两秒
	fmt.Println("时间", time.Now())
	t1 := <-timer.C //通道阻塞，直到时间达到了才执行
	fmt.Println(t1)

	stop := timer.Stop() //结束阻塞，直接运行下面的内容 当该运行到这个位置时计时器停止
	fmt.Println(stop)

	timer.Reset(time.Second * 2) //重新设置时间，修改newTimer时间

	ticker := time.NewTicker(time.Second)
	count := 1
	for _ = range ticker.C {
		fmt.Println("ticker...")
		count++
		if count > 5 {
			ticker.Stop()
			break
		}
	}
}
*/
