package utils

import (
    "image"
    "image/png"
    "log"
    "os"
)

func OutputToStdout(image image.Image) {
    err := png.Encode(os.Stdout, image)

    if err != nil {
        log.Fatal(err)
    }
}
