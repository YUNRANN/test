package main

import (
	"fmt"
	"time"
)

/*func aaa() {
	for i:=0;i<=100;i+=2{
		fmt.Println(i)
	}


	}

 go func (){
	for j:=1;j<=100;j+=2{
		fmt.Println(j)

	}()
}
*/

func main() {
	ch := make(chan int)
	var cn = make(chan int)
	go func() {

		for i := 0; i <= 100; i += 2 {

			fmt.Println(i)
			cn <- 1 //打印第一次i后阻塞此进程（1）  跳至进程二
			<-ch    //解锁阻塞进程二（4）

		}

	}()

	go func() {
		for j := 1; j <= 100; j += 2 {
			<-cn //解锁阻塞的进程一（2）
			fmt.Println(j)
			ch <- 2 //阻塞进程二（3）    调至进程一

		}

	}()
	time.Sleep(time.Second)
}
