package proj

import "C"
import (
	"unsafe"
)

func float64ToDoubleBytes(f float64) [8]byte {
	d := C.double(f)
	return *(*[8]byte)(unsafe.Pointer(&d))
}

func doubleBytesToFloat64(d []byte) float64 {
	var arr [8]byte
	copy(arr[:], d[:8])
	return float64(*(*C.double)(unsafe.Pointer(uintptr(unsafe.Pointer(&arr)))))
}
