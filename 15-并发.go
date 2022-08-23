package main

import (
	"fmt"
	"sync"
)

/*
 * 动态栈
 * 操作系统的线程一般都是固定的栈内存（通常为2MB）
 * GO中的 goroutine 非常的轻量级，初始栈空间很小（一般为2KB）
 * 所以，go 中一次创建数万个 goroutine也是可能的，而且go 的栈不是固定的，可以根据需求动态的扩容缩小
 * go 中的 runtime 会自动为goroutine 分配合适的栈空间
 */

/*
 * 调度
 * 操作系统内核在调度时会挂起当前正在执行的线程并将寄存器中的内容保存到内存中，然后选出接下来要执行的线程并从内存中恢复该线程的寄存器信息，然后恢复执行该线程的现场并开始执行线程。从一个线程切换到另一个线程需要完整的上下文切换。因为可能需要多次内存访问，索引这个切换上下文的操作开销较大，会增加运行的cpu周期。
 * goroutine 的调度是Go语言运行时（runtime）层面的实现，是完全由 Go 语言本身实现的一套调度系统——go scheduler。它的作用是按照一定的规则将所有的 goroutine 调度到操作系统线程上执行
 * Go 语言的调度器采用的是 GPM 调度模型
 *
 *
 * G：表示 goroutine，每执行一次go f()就创建一个 G，包含要执行的函数和上下文信息。
 * 全局队列（Global Queue）：存放等待运行的 G。
 * P：表示 goroutine 执行所需的资源，最多有 GOMAXPROCS 个。
 * P 的本地队列：同全局队列类似，存放的也是等待运行的G，存的数量有限，不超过256个。新建 G 时，G 优先加入到 P 的本地队列，如果本地队列满了会批量移动部分 G 到全局队列。
 * M：线程想运行任务就得获取 P，从 P 的本地队列获取 G，当 P 的本地队列为空时，M 也会尝试从全局队列或其他 P 的本地队列获取 G。M 运行 G，G 执行之后，M 会从 P 获取下一个 G，不断重复下去。
 * Goroutine 调度器和操作系统调度器是通过 M 结合起来的，每个 M 都代表了1个内核线程，操作系统调度器负责把内核线程分配到 CPU 的核上执行。
 *
 * 单从线程调度讲，Go语言相比起其他语言的优势在于OS线程是由OS内核来调度的，
 * goroutine 则是由Go运行时（runtime）自己的调度器调度的，完全是在用户态下完成的， 不涉及内核态与用户态之间的频繁切换，包括内存的分配与释放，都是在用户态维护着一块大的内存池，
 * 不直接调用系统的malloc函数（除非内存池需要改变），成本比调度OS线程低很多。 另一方面充分利用了多核的硬件资源，近似的把若干goroutine均分在物理线程上，
 * 再加上本身 goroutine 的超轻量级，以上种种特性保证了 goroutine 调度方面的性能。
 */

var wg sync.WaitGroup

func hello(i int) {
	defer wg.Done()
	fmt.Println("hello", i)
}
func main() {

	// wg.Add(1)
	// go hello()
	// fmt.Println("main is down")
	// wg.Wait()

	//启动多个goroutine
	//每次终端上打印数字的顺序都不一致。这是因为10个 goroutine 是并发执行的，而 goroutine 的调度是随机的。
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go hello(i)
	}
	fmt.Println("main is down")

	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println(i)
		}()
	}

	wg.Wait()
}
