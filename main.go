package gotools

import (
	"fmt"
	"unsafe"
)

// Simple print function used for debug
func DebugPrint(args ...interface{}) {
	if args[len(args)-1] != "\r\n" {
		args = append(args, "\r\n")
	}
	for _, v := range args {
		if v != "\r\n" {
			print(fmt.Sprint(v), " ")
		} else {
			print(fmt.Sprint(v))
		}

	}
}

// Not used, for read only
func toAddress(v interface{}) uintptr {
	return uintptr(unsafe.Pointer(&v))
}

// Not used
func toIntPointer(address uintptr) *int {
	return (*int)(unsafe.Pointer(address))
}

// Not used
func toInt(address uintptr) int {
	return *(*int)(unsafe.Pointer(address))
}
