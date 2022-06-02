package modes

import (
    "image"
    "log"

    "github.com/z3oxs/goshot/utils"
    "github.com/go-vgo/robotgo"
)

func Monitor(save string, clipboard, output, upload bool, display int) {
    var screenshot image.Image

    displays := utils.GetDisplays()

    if display > len(displays) {
        log.Fatal("Invalid screen selected.")

    }

    displaySelected := displays[display]

    screenshot = robotgo.CaptureImg(
        displaySelected.X,
        displaySelected.Y,
        displaySelected.Width,
        displaySelected.Height,
    )

    utils.PlaySound("/opt/goshot/screenshot.wav")

    if clipboard {
        utils.SaveToClipboard(screenshot)
    
    }

    if save != "" {
        robotgo.Save(screenshot, save)

    }

    if output {
        utils.OutputToStdout(screenshot)

    }

    if upload {
        utils.UploadImage(screenshot)

    }
}
