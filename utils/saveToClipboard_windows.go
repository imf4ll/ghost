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
    "os/exec"
    "strings"
    "bytes"
)

func SaveToClipboard(screenshot image.Image) {
    var image = new(bytes.Buffer)

    err = png.Encode(image, screenshot)
    if err != nil {
        log.Fatal(err)

    }

    imageToCString := C.CString(image.String())

    C.setClipboard(imageToCString)
}
