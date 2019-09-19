package util

import (
	"sync"
)

type FastQueue2 struct {
	data []interface{}
	fun func(interface{})
	sync.WaitGroup
}

func InitFastQueue2(worknum int,data []interface{},fun func(interface{}))  {
	lens:=len(data)
	workitemnum:=lens / worknum
	lastnum:=lens % worknum
	f:=&FastQueue2{
		data:data,
		fun:fun,
		WaitGroup:sync.WaitGroup{},
	}
	for i:=0;i<=worknum ;i++  {
		f.Add(1)
		start := i * workitemnum
		if i != worknum {
			go f.work( start , start + workitemnum )
			continue
		}
		go f.work( start , start + lastnum )
	}
	f.Wait()
	return
}
func (this *FastQueue2) work(start int, end int)  {
	defer this.Done()
	//end:=start+workitemnum
	for i:=start;i < end; i++ {
		this.fun(this.data[i])
		//fmt.Println("?")
	}
}