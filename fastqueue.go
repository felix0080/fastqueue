package fastqueue

import (
	"errors"
	"log"
	"math"
	"sync"
	"sync/atomic"
	"time"
)

/*
可写成一个并行处理库，
需要填充，处理方法，缓冲长度
*/
type FastQueue struct {
	c chan interface{}
	fun func(interface{})
	closec chan struct{}
	closesignal chan struct{} //阻塞的
	workNumNow int32
	workNumDesc int32
	sync.Once
}
func (this *FastQueue)close(){
	this.Do(func() {
		log.Println("已退出")
		close(this.closesignal)
	})
}
func BuildFastQueue(buflen, worknum int, method func(interface{})) *FastQueue {
	if buflen == 0 {
		buflen = 1000
	}
	fast := FastQueue{c:make(chan interface{},buflen),fun:method,closec:make(chan struct{},buflen),closesignal:make(chan struct{}),workNumDesc:int32(worknum)}
	for i := 0; i < worknum; i++ {
		//create work
		go fast.work()
	}
	return &fast
}
var timeOutErr = errors.New("timeout")
//可能会有延时，具体看消费速度
func (this *FastQueue) PutWithTimeout(item interface{},timeout time.Duration) error {
	select {
	case this.c <- item:
		return nil
	case <-time.After(timeout):
		return timeOutErr
	}
}
func (this *FastQueue) Put(item interface{})  {
	this.c <- item
}
func (this *FastQueue) ChangeWorkNum(z int) {
	atomic.AddInt32(&this.workNumDesc, int32(z))
	if z == 0 {
		return
	}
	if z > 0 {
		for i := 0; i < z ; i++ {
			//create work
			go this.work()
		}
	}
	if z < 0 {
		//send close
		fz:=math.Abs(float64(z))
		z=int(fz)
		for i := 0; i < z ; i++ {
			//close work
			this.closec<- struct{}{}
		}
	}
}
func(this *FastQueue) Close() {
	atomic.AddInt32(&this.workNumDesc, 0)
	close(this.c)
	<-this.closesignal
}
func (this *FastQueue) work() {
	atomic.AddInt32(&this.workNumNow, 1)
	//在此处对 总数进行统计，进来加一，退出减一，统计当前线程个数
	defer atomic.AddInt32(&this.workNumNow, -1)
	for  {
		select {
		case <-this.closec:
			//need close a worker,close self
			log.Println("已退出")
			return
		case c,ok:=<-this.c:
			if !ok {
				this.close()
				return
			}
			this.fun(c)
		}
	}
}

