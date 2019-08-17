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

