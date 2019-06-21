package gotools

import (
	"fmt"
	"unsafe"
)

var sizeOfInt = unsafe.Sizeof(0)

// Simple print function used for debug
func DebugPrint(args ...interface{}) {
	if args[len(args)-1] != "\r\n" {
		args = append(args, "\r\n")
	}

	for _, v := range args {
		switch v.(type) {
		case string:
			interfaceAddr := uintptr(unsafe.Pointer(&v))
			dataPtr := interfaceAddr + sizeOfInt
			dataAddr := *(*uintptr)(unsafe.Pointer(dataPtr)) // type:string

			dataAsSlice := *(*[]byte)(unsafe.Pointer(dataAddr))
			dataAsSliceLen := *(*int)(unsafe.Pointer(dataAddr + sizeOfInt))

			if !(dataAsSlice[dataAsSliceLen-1] == 10 && dataAsSlice[dataAsSliceLen-2] == 13) {
				print(fmt.Sprint(v), " ")
			} else {
				print(fmt.Sprint(v))
			}
		default:
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
