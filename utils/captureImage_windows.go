package utils

// go:build windows && !cgo

import (
    "image"

    "github.com/go-vgo/robotgo"
)

func CaptureScreen(filename string) image.Image {
    return robotgo.CaptureImg()

}

func CaptureRect(x, y, width, height int, filename string) image.Image {	
    return robotgo.CaptureImg(x, y, width, height)

}
