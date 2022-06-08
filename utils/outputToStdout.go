package utils

import (
    "image"
    "image/png"
    "image/jpeg"
    "log"
    "os"
)

func OutputToStdout(screenshot image.Image, format string) {
    var err error
    
    switch format {
        case "png":
            err = png.Encode(os.Stdout, screenshot)

        case "jpg":
            err = jpeg.Encode(os.Stdout, screenshot, &jpeg.Options {
                Quality: 100,
            })
    }

    if err != nil {
        log.Fatal(err)
    }
}
