package utils

import (
    "image"
    "image/png"
    "log"
    "os"
)

func OutputToStdout(screenshot image.Image) {
    err := png.Encode(os.Stdout, screenshot)
    if err != nil {
        log.Fatal(err)
    }
}
