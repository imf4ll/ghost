package utils

// go:build windows

/*
#cgo CFLAGS: -Wall

#include "../cgo/clipboard.c"
*/
import "C"

import (
    "log"
    "image"
    "image/png"
    "bytes"
)

func SaveToClipboard(screenshot image.Image) {
    var img = new(bytes.Buffer)

    err := png.Encode(img, screenshot)
    if err != nil {
        log.Fatal(err)

    }

    imageToCString := C.CString(img.String())

    C.setClipboard(imageToCString)
}
