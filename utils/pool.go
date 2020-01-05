package utils

import (
	"math"
	"sync"
	"unsafe"
)

// WaitGroupPool pool of WaitGroup
type WaitGroupPool struct {
	pool chan struct{}
	wg   *sync.WaitGroup
}

// hchan chan info
type hchan struct {
	qcount   uint           // total data in the queue
	dataqsiz uint           // size of the circular queue
	buf      unsafe.Pointer // points to an array of dataqsiz elements
	elemsize uint16
	closed   uint32
}

// NewWaitGroupPool creates a sized pool for WaitGroup
func NewWaitGroupPool(size int) *WaitGroupPool {
	if size <= 0 {
		size = math.MaxInt32
	}
	return &WaitGroupPool{
		pool: make(chan struct{}, size),
		wg:   &sync.WaitGroup{},
	}
}

// Add increments the WaitGroup counter by one.
// See sync.WaitGroup documentation for more information.
func (p *WaitGroupPool) Add() {
	p.pool <- struct{}{}
	p.wg.Add(1)
}

// Done decrements the WaitGroup counter by one.
// See sync.WaitGroup documentation for more information.
func (p *WaitGroupPool) Done() {
	<-p.pool
	p.wg.Done()
}

// Wait blocks until the WaitGroup counter is zero.
// See sync.WaitGroup documentation for more information.
func (p *WaitGroupPool) Wait() {
	p.wg.Wait()
}

// IsBlock
func (p *WaitGroupPool) IsBlock() bool {
	var c interface{} = p.pool
	i := (*[2]uintptr)(unsafe.Pointer(&c))
	h := (*hchan)(unsafe.Pointer(i[1]))
	return h.qcount >= h.dataqsiz
}

// IsClose
func (p *WaitGroupPool) IsClose() bool {
	var c interface{} = p.pool
	i := (*[2]uintptr)(unsafe.Pointer(&c))
	h := (*hchan)(unsafe.Pointer(i[1]))
	return h.closed == 1
}

// GetWorkCount
func (p *WaitGroupPool) GetWorkCount() uint {
	var c interface{} = p.pool
	i := (*[2]uintptr)(unsafe.Pointer(&c))
	h := (*hchan)(unsafe.Pointer(i[1]))
	return h.qcount
}

// GetQueueCount
func (p *WaitGroupPool) GetQueueCount() uint {
	var c interface{} = p.pool
	i := (*[2]uintptr)(unsafe.Pointer(&c))
	h := (*hchan)(unsafe.Pointer(i[1]))
	return h.dataqsiz
}
