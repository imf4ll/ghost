package utils

import (
    "image"
    "image/png"
    "image/jpeg"
    "log"
    "os"
)

func OutputToStdout(image image.Image, format string) {
    var err error

    switch format {
        case "png":
            err = png.Encode(os.Stdout, image)

        case "jpg":
            err = jpeg.Encode(os.Stdout, image, &jpeg.Options {
                Quality: 100,
            })
    }

    if err != nil {
        log.Fatal(err)
    }
}
