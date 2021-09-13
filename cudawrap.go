package main

// #cgo CFLAGS: -I.
// #cgo LDFLAGS: -L. -ltest
// #cgo LDFLAGS: -lcudart
// #cgo CXXFLAGS: -I/usr/lib/ -std=c++20
// #cgo LDFLAGS: -L/usr/lib/ -lstdc++ -O0
// #include <stdlib.h>
// #include "cuda-wrap.hpp"
import "C"
import (
	"fmt"
)

func exampleCuda(a, b int) {
	fmt.Println(
		"result via passing to cuda to add", a, b, C.test_add(C.int(a), C.int(b)),
	)
}
