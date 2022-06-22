package utils

// go:build linux

/*
#cgo LDFLAGS: -lX11

#include "../cgo/mouse.c"
*/
import "C"

import "unsafe"

type MousePosition struct {
    X1,
    X2,
    Y1,
    Y2 int
}

func GetMouse() MousePosition {
    var mousePos MousePosition

    getMousePos := (*[1 << 1]C.struct_MousePosition)(unsafe.Pointer(C.getMouse()))[0:1]

    m := getMousePos[0]

    mousePos.X1 = int(m.x1)
    mousePos.X2 = int(m.x2)
    mousePos.Y1 = int(m.y1)
    mousePos.Y2 = int(m.y2)

    return mousePos
}

func GetMouseAndKeyboard() MousePosition {
    var mousePos MousePosition

    getMousePos := (*[1 << 1]C.struct_MousePosition)(unsafe.Pointer(C.getMouseAndKeyboard()))[0:1]

    m := getMousePos[0]

    mousePos.X1 = int(m.x1)
    mousePos.X2 = int(m.x2)
    mousePos.Y1 = int(m.y1)
    mousePos.Y2 = int(m.y2)

    return mousePos
}
