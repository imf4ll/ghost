package utils

/*
#cgo CFLAGS: -Wall
#cgo LDFLAGS: -lXrandr -lX11

#include "../cgo/getDisplays.c"
*/
import "C"

import (
    "unsafe"
)

type Display struct {
    X int
    Y int
    Width int
    Height int
}

func GetDisplays() []Display {
    var validDisplays []Display

    displays := (*[1 << 10]C.struct_DisplayInfo)(unsafe.Pointer(C.getDisplays()))[0:10]

    for _, i := range displays {
        if (i.w != 0 && i.h != 0) {
            validDisplays = append(validDisplays, Display {
                X: int(i.x),
                Y: int(i.y),
                Width: int(i.w),
                Height: int(i.h),
            })

        }
    }

    return validDisplays
}
