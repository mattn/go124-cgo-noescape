# go124-cgo-noescape



## Benchmark

```
goos: windows
goarch: amd64
pkg: github.com/mattn/go124-cgo-noescape
cpu: AMD Ryzen 5 5600U with Radeon Graphics         
BenchmarkWithNoescape
BenchmarkWithNoescape-12       	24084874	        48.84 ns/op	       0 B/op	       0 allocs/op
BenchmarkWithoutNoescape
BenchmarkWithoutNoescape-12    	21355960	        54.42 ns/op	       4 B/op	       1 allocs/op
PASS
ok  	github.com/mattn/go124-cgo-noescape	2.673s
```

## License

MIT

## Author

Yasuhiro Matsumoto (a.k.a. mattn)
