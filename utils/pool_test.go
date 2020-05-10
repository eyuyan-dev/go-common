package utils_test

import (
	"fmt"
	"sync/atomic"
	"testing"
	"time"

	"github.com/gopkg-dev/go-common/utils"
)

func TestWaitGroupPool(t *testing.T) {
	wgp := utils.NewWaitGroupPool(10)

	var total uint32

	for i := 0; i < 100; i++ {
		wgp.Add()
		go func(total *uint32) {
			defer wgp.Done()
			atomic.AddUint32(total, 1)
			//t := ext.RangeRand(1, 5)
			//fmt.Println("Sleep ", t)
			time.Sleep(time.Second * time.Duration(1))
		}(&total)
		fmt.Println(wgp.IsBlock(), wgp.IsClose(), wgp.GetQueueCount(), wgp.GetWorkCount())
	}
	wgp.Wait()

	if total != 100 {
		t.Fatalf("The size '%d' of the pool did not meet expectations.", total)
	}
}
