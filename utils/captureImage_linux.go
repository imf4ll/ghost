package utils

// go:build linux

/*
#cgo CFLAGS: -Wall
#cgo LDFLAGS: -lX11 -lcairo

#include "../cgo/screenshot.c"
#include <cairo/cairo.h>
*/
import "C"

import (
    "os"
    "image"
    "log"
)

type CaptureType = C.cairo_surface_t

func CaptureScreen(filename string) image.Image {
    cFilename := C.CString(filename)
    
    C.captureScreen(cFilename)
    
    return getImage(filename)
}

func CaptureRect(x, y, width, height int, filename string) image.Image {
    cX := C.double(x)
    cY := C.double(y)
    cWidth := C.double(width)
    cHeight := C.double(height)
    cFilename := C.CString(filename)

    C.captureRect(cX, cY, cWidth, cHeight, cFilename)

    return getImage(filename)
}

func getImage(filename string) image.Image {
    var img image.Image
    var err error

    file, err := os.Open(filename)
    if err != nil {
        log.Fatal(err)

    }

    img, _, err = image.Decode(file)
    if err != nil {
        log.Fatal(err)

    }

    return img
}
