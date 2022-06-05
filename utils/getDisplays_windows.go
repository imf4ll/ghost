package utils

// go:build windows

/*
#cgo CFLAGS: -Wall
#cgo LDFLAGS: -lXrandr -lX11

#include "../cgo/getDisplays.c"
*/
import "C"

import "unsafe"

func GetDisplays() {

}
