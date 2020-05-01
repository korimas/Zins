package utils

import "unsafe"

func B2str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
