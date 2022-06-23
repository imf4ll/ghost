package utils

import (
    "os"
    "log"
    "image"
    "image/png"
)

func SaveImage(screenshot image.Image, save string) {
    file, err := os.Create(save)
    if err != nil {
        log.Fatal(err)
            
    }
    
    err = png.Encode(file, screenshot)
    if err != nil {
        log.Fatal(err)
    
    }
}
