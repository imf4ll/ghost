package utils

import "image"

func SaveHandler(clipboard, output, upload, file bool, screenshot image.Image) {
    if clipboard {
        SaveToClipboard(screenshot)
    
    }

    if output {
        OutputToStdout(screenshot)

    }

    if upload {
        UploadImage(screenshot)

    }

    if file {
        SaveImage(screenshot)

    }

}
