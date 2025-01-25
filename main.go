package main

/*
#include <stdio.h>
#include <stdint.h>

void do_write1(int32_t *v) {
	*v = 2;
}

void do_write2(int32_t *v) {
	*v = 3;
}

#cgo nocallback do_write1
#cgo noescape do_write1
*/
import "C"

//go:noinline
func DoWrite1() {
	var i C.int32_t = 1
	C.do_write1(&i)
}

//go:noinline
func DoWrite2() {
	var i C.int32_t = 1
	C.do_write2(&i)
}

func main() {
}
