## wyhash

A Go implementation of the 64-bit *wyhash* algorithm with a lot of optimizations. (final version 1)

original C++ implementation: https://github.com/wangyi-fudan/wyhash/blob/master/wyhash_final1.h



## QuickStart

```go
package main

import "github.com/zhangyunhao116/wyhash"

func main() {
	println(wyhash.Sum64Default([]byte("hello world!")))
}

```



## Benchmark

Go version: go1.15.6 linux/amd64

CPU: AMD 3700x(8C16T), running at 3.6GHz

OS: ubuntu 18.04

MEMORY: 16G x 2 (3200MHz)

```
name               time/op
Compare/0-16         3.92ns ± 0%
Compare/1-16         4.76ns ± 0%
Compare/2-16         4.76ns ± 0%
Compare/3-16         4.76ns ± 1%
Compare/4-16         4.32ns ± 1%
Compare/5-16         4.65ns ± 0%
Compare/6-16         4.72ns ± 4%
Compare/7-16         4.65ns ± 0%
Compare/8-16         4.40ns ± 0%
Compare/9-16         4.69ns ± 0%
Compare/10-16        4.69ns ± 0%
Compare/11-16        4.69ns ± 0%
Compare/12-16        4.71ns ± 1%
Compare/13-16        4.69ns ± 0%
Compare/14-16        4.69ns ± 1%
Compare/15-16        4.73ns ± 2%
Compare/16-16        4.73ns ± 3%
Compare/17-16        5.49ns ± 0%
Compare/21-16        5.48ns ± 0%
Compare/24-16        5.17ns ± 4%
Compare/29-16        5.56ns ± 4%
Compare/32-16        5.46ns ± 0%
Compare/33-16        6.61ns ± 0%
Compare/64-16        7.70ns ± 0%
Compare/69-16        7.99ns ± 0%
Compare/96-16        9.03ns ± 0%
Compare/97-16        9.75ns ± 0%
Compare/128-16       10.8ns ± 0%
Compare/129-16       10.9ns ± 0%
Compare/240-16       15.9ns ± 1%
Compare/241-16       16.8ns ± 1%
Compare/512-16       28.4ns ± 1%
Compare/1024-16      50.5ns ± 0%
Compare/102400-16    4.26µs ± 1%

name               speed
Compare/0-16
Compare/1-16        210MB/s ± 0%
Compare/2-16        420MB/s ± 0%
Compare/3-16        631MB/s ± 0%
Compare/4-16        925MB/s ± 0%
Compare/5-16       1.08GB/s ± 0%
Compare/6-16       1.27GB/s ± 4%
Compare/7-16       1.51GB/s ± 0%
Compare/8-16       1.82GB/s ± 0%
Compare/9-16       1.92GB/s ± 0%
Compare/10-16      2.13GB/s ± 0%
Compare/11-16      2.34GB/s ± 0%
Compare/12-16      2.55GB/s ± 1%
Compare/13-16      2.77GB/s ± 0%
Compare/14-16      2.98GB/s ± 2%
Compare/15-16      3.16GB/s ± 4%
Compare/16-16      3.39GB/s ± 3%
Compare/17-16      3.10GB/s ± 0%
Compare/21-16      3.82GB/s ± 2%
Compare/24-16      4.65GB/s ± 4%
Compare/29-16      5.22GB/s ± 4%
Compare/32-16      5.86GB/s ± 0%
Compare/33-16      4.99GB/s ± 0%
Compare/64-16      8.31GB/s ± 0%
Compare/69-16      8.64GB/s ± 0%
Compare/96-16      10.6GB/s ± 0%
Compare/97-16      9.95GB/s ± 0%
Compare/128-16     11.9GB/s ± 0%
Compare/129-16     11.8GB/s ± 0%
Compare/240-16     15.1GB/s ± 1%
Compare/241-16     14.4GB/s ± 1%
Compare/512-16     18.0GB/s ± 1%
Compare/1024-16    20.3GB/s ± 0%
Compare/102400-16  24.0GB/s ± 1%
```

