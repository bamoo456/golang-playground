
#### Basic concept
In the golang, we can use package `runtime/pprof` in our code to profile the program.
> e.g.
```
func test() {
    pprof.StartCPUProfile(f)
    ...
    ...
    
    defer pprof.StopCPUProfile()
}

```

The other way is by using the benchmark functions, in this case the Profiling 
is part of benchmark tests with some flags specified 

`go test -bench=...`

 

##### example of profiling in tests
```
# run the benchmark function `BenchmarkCalcAvg` and output its cpu profile log 
go test -run=NONE -bench=BenchmarkCalcAvg -cpuprofile=cpu.log

# run the benchmark function `ToJSONArray` of the package `golang-playground/practices/profiling` and ouput its log
go test -run=NONE -bench=ToJSONArray -cpuprofile=profile.log golang-playground/practices/profiling

# parse the log file
go tool pprof -text -nodecount={how many significant operations} {the profile binary file} {the profile log file}

# Use the go tool the parse the prifile log file with text output
go tool pprof -text -nodecount=10 ./lib.test profile.log

# Use the go tool the parse the prifile log file with pdf output
go tool pprof -pdf -nodecount=10 ./lib.test profile.log > tmp.pdf
```

> NOTE: the profile binary file usually named as {package name}.test (e.g.: lib.test)

<br>
<br>
<br>

##### Another good tool for cpu profiling - flame graph 
https://github.com/uber/go-torch

*basic usage:*

```
# install the go-torch
go get github.com/uber/go-torch`

# cd to the target package
git clone https://github.com/brendangregg/FlameGraph.git

# generate the flame graph file
go-torch profiling.test profile.log

```


#### Result Analysis - CPU
In the cpu profile text output, the frequently used function will be listed in the first place.
In the ToJSONArray function profiling report, we can know the program spends much time on 
functions `runtime.mach_semaphore_signal` and `runtime.mallocgc`(GC) . 

This means we may able to optimize the `ToJSONArray` by reducing the frequency of allocate new memory.
> NOTE: just an example on how to profiling the function 
```
820ms of 870ms total (94.25%)
Showing top 10 nodes out of 51 (cum >= 10ms)
      flat  flat%   sum%        cum   cum%
     350ms 40.23% 40.23%      350ms 40.23%  runtime.mach_semaphore_signal
     200ms 22.99% 63.22%      670ms 77.01%  runtime.mallocgc
      70ms  8.05% 71.26%       70ms  8.05%  runtime.heapBitsSetType
      60ms  6.90% 78.16%      740ms 85.06%  golang-playground/practices/profiling.ToJSONArray
      40ms  4.60% 82.76%      780ms 89.66%  golang-playground/practices/profiling.BenchmarkToJSONArray
      40ms  4.60% 87.36%       40ms  4.60%  runtime.mach_semaphore_wait
      20ms  2.30% 89.66%       20ms  2.30%  runtime.(*mspan).refillAllocCache
      20ms  2.30% 91.95%       20ms  2.30%  runtime.mach_semaphore_timedwait
      10ms  1.15% 93.10%       20ms  2.30%  runtime.(*mcache).refill
      10ms  1.15% 94.25%       10ms  1.15%  runtime.(*mheap).alloc_m

```


#### Result Analysis - Memory
In the memory profile text output, the frequently used function will be listed in the first place.

```
0 of 0 total (    0%)
      flat  flat%   sum%        cum   cum%
         0     0%     0%          0     0%  golang-playground/practices/profiling.BenchmarkToJSONArray
         0     0%     0%          0     0%  golang-playground/practices/profiling.ToJSONArray
         0     0%     0%          0     0%  runtime.goexit
         0     0%     0%          0     0%  testing.(*B).launch
         0     0%     0%          0     0%  testing.(*B).runN

```


### Runtime profiling
Basically, the runtime profiling is by usgng package `runtime/pprof` or `net/http/pprof`

#### CPU profiling
To do the cpu profiling in runtime, we can refer below two packages for reference:
- `https://golang.org/pkg/runtime/pprof/`
> NOTE: I can only generate empty profile output by this method currently, need more time to dig into this.

- `https://github.com/pkg/profile`
> NOTE: I have never tried it.


#### Memory profiling
With `net/http/pprof` enable, we can have an http interface on the program we want, and check the runtime information
of the program. It's good for checking the memory leak issue.

>NOTE: check the `server.go` on how to use it.

##### Directly access the interface
- use browser or any http client to access `http://localhost:6060/debug/pprof/`

##### Directly access the interface for heap information
- use browser or any http client to access `http://localhost:6060/debug/pprof/heap?debug=1`

##### Use go tool pprof to access the interface
- use command here `go tool pprof http://localhost:6060/debug/pprof/heap` to get the information of heap
> NOTE: we can also use `go-torch` to debug it `go-torch -inuse_space http://127.0.0.1:6060/debug/pprof/heap --colors=mem`


REFERENCE:
http://fuxiaohei.me/2015/10/14/pugo-mem-leak-profile.html
http://www.philo.top/2015/05/29/golangProfilingAndGC2/
https://lrita.github.io/2017/05/26/golang-memory-pprof/
http://cjting.me/golang/use-pprof-to-optimize-go/
http://xiaorui.cc/2016/03/20/golang%E4%BD%BF%E7%94%A8pprof%E7%9B%91%E6%8E%A7%E6%80%A7%E8%83%BD%E5%8F%8Agc%E8%B0%83%E4%BC%98/
