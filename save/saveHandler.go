package save

import (
    "image"

    "github.com/imf4ll/ghost/utils"
)

func SaveHandler(clipboard, output, upload, file bool, screenshot image.Image) {
    if clipboard {
        SaveToClipboard(screenshot)
    
    }

    if output {
        utils.OutputToStdout(screenshot)

    }

    if upload {
        utils.UploadImage(screenshot)

    }

    if file {
        SaveImage(screenshot)

    }

}
