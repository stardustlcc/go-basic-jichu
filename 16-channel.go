package main

import "fmt"

/*
 * Go语言采用的并发模型是CSP（Communicating Sequential Processes），提倡通过通信共享内存而不是通过共享内存而实现通信。
 * 如果说 goroutine 是Go程序并发的执行体，channel就是它们之间的连接。channel是可以让一个 goroutine 发送特定值到另一个 goroutine 的通信机制。
 * Go 语言中的通道（channel）是一种特殊的类型。通道像一个传送带或者队列，总是遵循先入先出（First In First Out）的规则，保证收发数据的顺序。每一个通道都是一个具体类型的导管，也就是声明channel的时候需要为其指定元素类型。
 */

/*
 * 声明channel
 * var 变量名称 chan 元素类型
 * chan ：是关键字
 * 元素类型：是指通道中传递元素的类型
 *
 * 注意：没有初始化的通道类型变量默认零值是nil
 */

/*
 	关闭后的通道有以下特点：
	对一个关闭的通道再发送值就会导致 panic。
	对一个关闭的通道进行接收会一直获取值直到通道为空。
	对一个关闭的并且没有值的通道执行接收操作会得到对应类型的零值。
	关闭一个已经关闭的通道会导致 panic。
*/

/*
 * 无缓冲的通道
 * 无缓冲的通道又称为阻塞的通道
 * 我们使用ch := make(chan int)创建的是无缓冲的通道，无缓冲的通道只有在有接收方能够接收值的时候才能发送成功，否则会一直处于等待发送的阶段。
 * 同理，如果对一个无缓冲通道执行接收操作时，没有任何向通道中发送值的操作那么也会导致接收操作阻塞。
 */

/*
 * 单向通道
 * 单向通道可以作为参数在多个任务函数间进行传递，用来约束参数是只读，还是只存的通道
 * <- chan int // 只接收通道，只能接收不能发送
 * chan <- int // 只发送通道，只能发送不能接收
 */
func main() {

	//没有初始化的通道类型变量默认零值是nil
	var test chan int
	fmt.Println(test) //nil

	//初始化channel
	test = make(chan int, 1)
	fmt.Println(test) //0xc00008c060

	//有缓冲区大小，可选
	//test = make(chan int, 10)

	//发送一个值到通道中
	test <- 10
	x := <-test
	//忽略值
	//<-test
	fmt.Println(x)

	//关闭通道
	close(test)

	/* -------------------------*/

	//无缓冲通道案例
	ch := make(chan int)
	//开启一个协程，静候接收通道值
	go getCh(ch)
	//像无缓冲的通道发送值，如果没有人接则会阻塞
	ch <- 10
	fmt.Println("发送成功")

	//有缓冲通道
	ch1 := make(chan int, 1)
	ch1 <- 100
	fmt.Println("有缓冲通道发送成功")

	//多返回值模式
	//value：从通道中取出的值，如果通道被关闭则返回对应类型的零值
	//ok：通道ch关闭时返回 false，否则返回 true。
	//value, ok := <- ch1

	ch2 := make(chan int, 2)
	ch2 <- 10
	ch2 <- 20
	close(ch2)
	//fun1(ch2)

	//for range 接收值
	//当通道被关闭后，会在通道内的所有值被接收完毕后会自动退出循环
	fun2(ch2)

	/* -------------------------*/

}

func getCh(ch chan int) {
	fmt.Println("我接到数据了：", <-ch)
}

func fun1(ch chan int) {
	for {
		v, ok := <-ch
		if ok == false {
			fmt.Println("通道已关闭")
			break
		}
		fmt.Printf("v:%#v ok:%#v\n", v, ok)
	}
}

//for range 接收值
//当通道被关闭后，会在通道内的所有值被接收完毕后会自动退出循环
func fun2(ch chan int) {
	for v := range ch {
		fmt.Println(v)
	}
}
