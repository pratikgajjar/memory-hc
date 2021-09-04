# memory-hc
GoLang healthcheck server based on memory usuage


```bash
➜  memory-hc git:(develop) ✗ ./memory-hc -h
Usage of ./memory-hc:
  -cli
    	Cli only
  -http.addr string
    	HTTP listen address (default ":8080")
  -mem.threshold float
    	Healthy Memory Threshold (default 75)
```

## Usage 
- Situations where you can't identify memory leak and want to take action based on current memory usage of machine

## Example
```shell
./memory-hc -mem.threshold 86 -cli && echo "Hello"
```
This will print hello only if system mem usage is below 86%
