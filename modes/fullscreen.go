package modes

import (
    "github.com/go-vgo/robotgo"
    "github.com/z3oxs/ghost/utils"
)

func Fullscreen(save string, clipboard, output, upload bool) {
    screenshot := robotgo.CaptureImg()

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
