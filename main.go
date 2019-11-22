package main

/*
#include "lib.h"
*/
import "C"
import (
	"fmt"
	"unsafe"
)

//export GetData
func GetData(gm *C.GoMem) {
	str := "Golang"
	gm.data = unsafe.Pointer(C.CString(str))
	gm.size = C.int(len(str))
}

//export SetData
func SetData(gm *C.GoMem) {
	var str string
	if gm.size == 0 {
		str = C.GoString((*C.char)(gm.data))
	} else {
		str = C.GoStringN((*C.char)(gm.data), gm.size)
	}
	fmt.Println("SetData:", str)
}
func main() {
	var gm C.GoMem
	C.GetData(&gm) 
	fmt.Println("return:", gm.data, gm.size) 
	data := make([]byte, gm.size) 
	C.memcpy(unsafe.Pointer(&data[0]), gm.data, C.ulonglong(gm.size)) 
	fmt.Println(string(data))
}
