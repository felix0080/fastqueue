package util

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"testing"
	"time"
)
//对此脚本进行pprof 看看为什么这么慢，主要是什么原因导致的
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
