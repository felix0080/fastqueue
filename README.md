
# 本库的目的是：将数组中元素分发到多个work进行处理

# 采用FastQueue方案，将item使用chan分发给多个work处理，此方案经过测试，绝大多数时间都浪费在乐lock上，且资源占有太大，需要再次思索chan的使用场景，和如何更优化的使用chan


# 采用FastQueue2方案，使用并发式读取，免去竞争的性能可以媲美直接遍历，在有工作负荷时性能远超单线程遍历。
将数组中元素分发到多个work进行处理
   goos: windows  
   goarch: amd64  
   pkg: app/util  
   BenchmarkName-4   	     100	  18319506 ns/op  
   PASS  

损耗不大，可采用全并发，不锁的方案。  
   goos: windows  
   goarch: amd64  
   pkg: app/util  
   BenchmarkName-4   	     300	   4797241 ns/op  
   PASS  
