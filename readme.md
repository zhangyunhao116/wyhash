## wyhash

A Go implementation of the 64-bit *wyhash* algorithm with a lot of optimizations. (final version 1)

original C++ implementation: https://github.com/wangyi-fudan/wyhash/blob/master/wyhash_final1.h



## QuickStart

```go
package main

import "github.com/zhangyunhao116/wyhash"

func main() {
	println(wyhash.Sum64String("hello world!"))
}

```



## Benchmark

Go version: go1.15.6 linux/amd64

CPU: AMD 3700x(8C16T), running at 3.6GHz

OS: ubuntu 18.04

MEMORY: 16G x 2 (3200MHz)

```
name              time/op
Wyhash/0-16         3.08ns ± 0%
Wyhash/1-16         4.09ns ± 0%
Wyhash/2-16         4.09ns ± 0%
Wyhash/3-16         4.09ns ± 1%
Wyhash/4-16         3.29ns ± 2%
Wyhash/5-16         3.71ns ± 1%
Wyhash/6-16         3.71ns ± 0%
Wyhash/7-16         3.76ns ± 3%
Wyhash/8-16         3.39ns ± 2%
Wyhash/9-16         3.78ns ± 2%
Wyhash/10-16        3.76ns ± 0%
Wyhash/11-16        3.76ns ± 0%
Wyhash/12-16        3.78ns ± 1%
Wyhash/13-16        3.76ns ± 1%
Wyhash/14-16        3.79ns ± 2%
Wyhash/15-16        3.77ns ± 0%
Wyhash/16-16        3.76ns ± 0%
Wyhash/17-16        4.98ns ± 0%
Wyhash/21-16        4.79ns ± 0%
Wyhash/24-16        4.39ns ± 2%
Wyhash/29-16        4.81ns ± 1%
Wyhash/32-16        4.77ns ± 0%
Wyhash/33-16        6.04ns ± 0%
Wyhash/64-16        7.25ns ± 1%
Wyhash/69-16        7.33ns ± 0%
Wyhash/96-16        8.31ns ± 0%
Wyhash/97-16        9.19ns ± 0%
Wyhash/128-16       10.4ns ± 0%
Wyhash/129-16       10.2ns ± 0%
Wyhash/240-16       15.7ns ± 3%
Wyhash/241-16       16.7ns ± 1%
Wyhash/512-16       28.2ns ± 1%
Wyhash/1024-16      50.2ns ± 1%
Wyhash/102400-16    4.27µs ± 0%

name              speed
Wyhash/0-16
Wyhash/1-16        245MB/s ± 0%
Wyhash/2-16        489MB/s ± 0%
Wyhash/3-16        733MB/s ± 1%
Wyhash/4-16       1.22GB/s ± 2%
Wyhash/5-16       1.35GB/s ± 1%
Wyhash/6-16       1.62GB/s ± 0%
Wyhash/7-16       1.86GB/s ± 3%
Wyhash/8-16       2.36GB/s ± 2%
Wyhash/9-16       2.39GB/s ± 0%
Wyhash/10-16      2.66GB/s ± 0%
Wyhash/11-16      2.92GB/s ± 0%
Wyhash/12-16      3.17GB/s ± 1%
Wyhash/13-16      3.45GB/s ± 0%
Wyhash/14-16      3.69GB/s ± 2%
Wyhash/15-16      3.99GB/s ± 0%
Wyhash/16-16      4.26GB/s ± 0%
Wyhash/17-16      3.41GB/s ± 0%
Wyhash/21-16      4.39GB/s ± 0%
Wyhash/24-16      5.46GB/s ± 2%
Wyhash/29-16      6.03GB/s ± 1%
Wyhash/32-16      6.72GB/s ± 0%
Wyhash/33-16      5.47GB/s ± 0%
Wyhash/64-16      8.83GB/s ± 1%
Wyhash/69-16      9.42GB/s ± 0%
Wyhash/96-16      11.6GB/s ± 0%
Wyhash/97-16      10.6GB/s ± 0%
Wyhash/128-16     12.3GB/s ± 0%
Wyhash/129-16     12.6GB/s ± 0%
Wyhash/240-16     15.3GB/s ± 2%
Wyhash/241-16     14.5GB/s ± 1%
Wyhash/512-16     18.1GB/s ± 1%
Wyhash/1024-16    20.4GB/s ± 1%
Wyhash/102400-16  24.0GB/s ± 0%
```

