package utils

/*
#cgo CFLAGS: -Wall
#cgo LDFLAGS: 

#include "../cgo/clipboard.c"
*/
import "C"

import (
    "log"
    "image"
    "os/exec"
    "runtime"
    "strings"

    "github.com/go-vgo/robotgo"
)

func SaveToClipboard(screenshot image.Image) {
    if runtime.GOOS == "linux" {
        whereIsXClip, err := exec.Command("which", "xclip").Output()
        if err != nil {
            log.Fatal(err)

        }
        
        if strings.Contains(string(whereIsXClip), "not found") {
            log.Fatal("XClip was not found, are you sure that you have installed?")

        }

        robotgo.Save(screenshot, "/tmp/screenshot.png")

        cmd := exec.Command("xclip", "-sel", "clipboard", "-t", "image/png", "-i", "/tmp/screenshot.png")

        if err := cmd.Run(); err != nil {
            log.Fatal(err)

        }

    } else if runtime.GOOS == "windows" {

    }
}
