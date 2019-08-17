# gologger
_Gologger_ is a simple _Go_ logging library, focusing on _asynchronous_ and _daily rotate file_.

## Installation
```
$ go get github.com/nikoeleison/gologger
```

## Example
### Code
```go
package main

import "github.com/nikoeleison/gologger"

func main() {
	logger := gologger.New(
		"./log",
		"gologger",
	)
	//file path
	//file name
	//file output: ./log/gologger.log.2006-01-02

	defer logger.Kill()
	//flush te rest of log pool
	//close gologger

	//Info flavour
	logger.Info("Hello World!")
	logger.Infof("%s" "Hello World!")
	//Error flavour
	logger.Error("Hello World!")
	logger.Errorf("%s" "Hello World!")
	//Debug flavour
	logger.Debug("Hello World!")
	logger.Debugf("%s" "Hello World!")
	//Panic flavour
	logger.Panic("Hello World!")
	logger.Panicf("%s" "Hello World!")
}
```

### Output:
```
2006-01-02 15:04:05.000 [ INFO] main.go:19                     - Hello World!
2006-01-02 15:04:05.000 [ INFO] main.go:20                     - Hello World!
2006-01-02 15:04:05.000 [ERROR] main.go:22                     - Hello World!
2006-01-02 15:04:05.000 [ERROR] main.go:23                     - Hello World!
2006-01-02 15:04:05.000 [DEBUG] main.go:25                     - Hello World!
2006-01-02 15:04:05.000 [DEBUG] main.go:26                     - Hello World!
2006-01-02 15:04:05.000 [PANIC] main.go:28                     - Hello World!
2006-01-02 15:04:05.000 [PANIC] main.go:29                     - Hello World!
```

## Benchmark
* Date: 17 August 2019
* Memory: 3,6 GiB
* Processor: Intel® Core™ i3-6006U CPU @ 2.00GHz × 4

### Go benchmark
`$ go test -bench=. -benchtime=5s`

| test | times ran | ns/op | B/op | allocs/op |
|------|-----------|-------|------|-----------|
| **BenchmarkGologger** | 10000 | 765113 | 8306 | 143 |

### Apache benchmark
`$ ab -n 1000 -c 1000 localhost:3000/index`

```
This is ApacheBench, Version 2.3 <$Revision: 1706008 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 100 requests
Completed 200 requests
Completed 300 requests
Completed 400 requests
Completed 500 requests
Completed 600 requests
Completed 700 requests
Completed 800 requests
Completed 900 requests
Completed 1000 requests
Finished 1000 requests


Server Software:
Server Hostname:       localhost
Server Port:           3000

Document Path:         /index
Document Length:       8 bytes

Concurrency Level:     1000
Time taken for tests:  0.175 seconds
Complete requests:     1000
Failed requests:       991
   (Connect: 0, Receive: 0, Length: 991, Exceptions: 0)
Total transferred:     126794 bytes
HTML transferred:      9893 bytes
Requests per second:   5715.07 [#/sec] (mean)
Time per request:      174.976 [ms] (mean)
Time per request:      0.175 [ms] (mean, across all concurrent requests)
Transfer rate:         707.65 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    5   9.6      1      26
Processing:     1   18   7.2     17      47
Waiting:        1   18   7.3     16      47
Total:          1   23  11.9     20      72

Percentage of the requests served within a certain time (ms)
  50%     20
  66%     21
  75%     30
  80%     33
  90%     41
  95%     43
  98%     59
  99%     61
 100%     72 (longest request)
```
