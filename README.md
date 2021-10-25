# memory-hc
GoLang health check server based on memory usage


```shell
➜  memory-hc git:(develop) ✗ ./memory-hc -h
Usage of bin/memory_hc:
  -cli
    	Cli Only
  -http.addr string
    	HTTP listen address (default ":8080")
  -mem.threshold float
    	Healthy Memory Threshold (default 75)
  -unhealthy
    	Exit with code 0 on unhealthy
```

## Usage 
- Situations where you can't identify memory leak and want to take action based on current memory usage of machine

## Example
```shell
➜  memory-hc git:(develop) ✗ bin/memory_hc -cli -unhealthy -mem.threshold 89 ; echo $?
2021/10/25 17:53:34 Mem Per% 85.90097419437767 Mem Used 14754181120 Mem Total 17175801856 Mem Free 65474560
1
➜  memory-hc git:(develop) ✗ bin/memory_hc -cli -unhealthy -mem.threshold 50 ; echo $?
2021/10/25 17:53:38 Mem Per% 86.02720328249858 Mem Used 14776143872 Mem Total 17176129536 Mem Free 43651072
0
➜  memory-hc git:(develop) ✗ bin/memory_hc -cli  -mem.threshold 50 ; echo $?
2021/10/25 17:53:44 Mem Per% 85.9415910515156 Mem Used 14762188800 Mem Total 17177001984 Mem Free 58572800
1
➜  memory-hc git:(develop) ✗ bin/memory_hc -cli  -mem.threshold 90 ; echo $?
2021/10/25 17:53:50 Mem Per% 85.89398383289219 Mem Used 14754349056 Mem Total 17177395200 Mem Free 66715648
0
➜  memory-hc git:(develop) bin/memory_hc
2021/10/25 17:57:10 Starting HTTP server on :8080
```
