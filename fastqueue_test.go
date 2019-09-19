package fastqueue

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"testing"
	"time"
)
//对此脚本进行pprof 看看为什么这么慢，主要是什么原因导致的
//猜测是 线程内部G队列等待数目较多，导致锁时间过长
//解决方法，还是尽量采用无chan的方式，定义一个可以分批进行for循环的方法即可
func TestFast(t *testing.T) {
	//60-82 s
	//远程获取pprof数据
	go func() {
		log.Println(http.ListenAndServe("localhost:8080", nil))
	}()
	slices:=make([]string,100000000)
	r:=int32(0)
	fast:=BuildFastQueue(100000,1, func(i interface{}) {
		/*_,ok := i.(string)
		if ok {
			//fmt.Println(".")
			atomic.AddInt32(&r,1)
		}*/
	})
	for i:=0;i< 100000000; i++ {
		fast.Put(slices[i])
	}

	fast.Close()
	fmt.Println(r)
	/*for i:=0;i< 100000000; i++ {
		//fmt.Println(".")
		s:=string(slices[i])
		if s == "" {

		}
		atomic.AddInt32(&r,1)
	}
	fmt.Println(r)*/
	time.Sleep(time.Minute)
}
func TestOIUY(t *testing.T) {
	go func() {
		log.Println(http.ListenAndServe("localhost:8080", nil))
	}()

	for i := 0 ; i < 10 ; i++  {
		fmt.Println(len(slice))
		InitFastQueue2(4,slice, func(i interface{}) {
			//fmt.Print()
		})
		//time.Sleep(1*time.Second)
		fmt.Println("finish")
	}
}
func init() {
	for i:=0 ; i < 12345678 ; i++  {
		slice= append(slice, i)
	}
}
var slice []interface{}
func BenchmarkName(b *testing.B) {
	/*
	      goos: windows
	      goarch: amd64
	      pkg: app/util
	      BenchmarkName-4   	     100	  18319506 ns/op
	      PASS
	   损耗不大，可采用全并发，不锁的方案。
	*/
	/*
	   goos: windows
	   goarch: amd64
	   pkg: app/util
	   BenchmarkName-4   	     300	   4797241 ns/op
	   PASS
	*/
	for i := 0; i < b.N; i++ {
		//InitFastQueue2(4,slice, func(i interface{}) {
		//	//fmt.Println(i)
		//})
		for i := 0; i < 12345678; i++ {

		}
	}
}