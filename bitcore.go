package bitcore_hash

// #cgo CFLAGS: -I.
// #cgo LDFLAGS: -L. -lbitcore_hash
// #include "bitcore.h"
import "C"

import (
	"unsafe"
)

func GetPowHash(hash []byte, timestamp int32) []byte {
	powhash := make([]byte, 32)
	C.bitcore_hash((*C.char)(unsafe.Pointer(&hash[0])), (*C.char)(unsafe.Pointer(&powhash[0])), (C.uint32_t)(timestamp))
	return powhash
}
