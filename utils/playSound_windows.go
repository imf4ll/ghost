package utils

// go:build windows

/*
#cgo LDFLAGS: -lwinmm

#include "../cgo/sound.c"
*/
import "C"

func PlaySound() {
    C.play()

}
