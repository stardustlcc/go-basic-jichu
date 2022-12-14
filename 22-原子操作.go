package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

/*
 * 针对整数数据类型（int32、uint32、int64、uint64）我们还可以使用原子操作来保证并发安全，
 * 通常直接使用原子操作比使用锁操作效率更高。Go语言中原子操作由内置的标准库sync/atomic提供
 * 读操作 LoadInt32
 * 写操作 StoreInt32
 * 修改操作 AddInt32
 * 交换操作 SwapInt32
 * 比较并交换操作  CompareAndSwapInt32
 */

type Counter interface {
	Inc()
	Load() int64
}

//普通版本-----------------------
type CommonCounter struct {
	counter int64
}

func (c CommonCounter) Inc() {
	c.counter++
}

func (c CommonCounter) Load() int64 {
	return c.counter
}

//互斥锁版本-----------------------
type MutexCounter struct {
	counter int64
	lock    sync.Mutex
}

func (m *MutexCounter) Inc() {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.counter++
}

func (m *MutexCounter) Load() int64 {
	m.lock.Lock()
	defer m.lock.Unlock()
	return m.counter
}

//原子操作版本-----------------------
type AtomicCounter struct {
	counter int64
}

func (a *AtomicCounter) Inc() {
	atomic.AddInt64(&a.counter, 1)
}

func (a *AtomicCounter) Load() int64 {
	return atomic.LoadInt64(&a.counter)
}

func main() {

	// 非并发安全案例
	c1 := CommonCounter{}
	test(c1)

	// 使用互斥锁实现并发安全
	c2 := MutexCounter{}
	test(&c2)

	// 并发安全且比互斥锁效率更高
	c3 := AtomicCounter{}
	test(&c3)
}

func test(c Counter) {
	var wg sync.WaitGroup
	start := time.Now()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			c.Inc()
			wg.Done()
		}()
	}
	wg.Wait()
	end := time.Now()
	fmt.Println(c.Load(), end.Sub(start))
}
