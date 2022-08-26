package main

import (
	"fmt"
	"sync"
	"time"
)

/*
 * 读写锁
 * 互斥锁是完全互斥的，实际很多场景都是读多写少，当我们并发的读取一个资源而不涉及资源修改是没必要加互斥锁的。
 * 这种场景使用读写锁是更好的一种选择。
 * 读写锁在go语言中使用 sync 包中的 RWMute类型
 * sync.RWMute 提供了以下5个方法
 * 1）Lock() 获取写锁
 * 2）Unlock() 释放写锁
 * 3）Rlock() 获取读锁
 * 4）RUnlock() 释放读锁
 * 5）RLocker() Locker 返回一个实现Locker接口的读写锁
 *
 * 读写锁分为两种：读锁和写锁
 * 当一个 goroutine 获取到读锁之后，其他的 goroutine 如果是获取读锁会继续获得锁，如果是获取写锁就会等待
 * 而当一个 goroutine 获取写锁之后，其他的 goroutine 无论是获取读锁还是写锁都会等待
 */

var (
	x       int64
	wg      sync.WaitGroup
	mutex   sync.Mutex
	rwMutex sync.RWMutex
)

//使用互斥锁的写操作
func writeLock() {
	mutex.Lock()
	x = x + 1
	time.Sleep(10 * time.Millisecond)
	mutex.Unlock()
	wg.Done()
}

//使用互斥锁的读操作
func readLock() {
	mutex.Lock()
	time.Sleep(time.Millisecond) //1毫秒
	mutex.Unlock()
	wg.Done()
}

//读写互斥锁-写操作
func writeRWLock() {
	rwMutex.Lock() //加写锁
	x = x + 1
	time.Sleep(10 * time.Millisecond)
	rwMutex.Unlock() //释放写锁
	wg.Done()
}

//读写互斥锁-读操作
func readRWLock() {
	rwMutex.RLock() //加读锁
	time.Sleep(time.Millisecond)
	rwMutex.RUnlock() //释放读锁
	wg.Done()
}

func do(wf, rf func(), wc, rc int) {
	start := time.Now()
	//wc 并发写操作
	for i := 0; i < wc; i++ {
		wg.Add(1)
		go wf()
	}

	//rc 个并发读操作
	for i := 0; i < rc; i++ {
		wg.Add(1)
		go rf()
	}
	wg.Wait()
	con := time.Since(start)
	fmt.Printf("x:%v con:%v\n", x, con)
}

func main() {

	// 使用互斥锁，10并发写，1000并发读
	do(writeLock, readLock, 10, 1000)

	// 使用读写互斥锁，10并发写，1000并发读
	do(writeRWLock, readRWLock, 10, 1000)
}
