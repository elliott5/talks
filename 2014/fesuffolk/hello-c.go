// +build OMIT

package main

/*
#include "strings.h"
*/
import "C"
import "fmt"

func main() {
	hello := "Hello 世界"
	length := C.strlen(C.CString(hello))
	fmt.Printf("%s (len in C: %d, and in Go: %d)\n", hello, length, len(hello))
}
