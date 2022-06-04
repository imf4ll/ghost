package modes

import (
    "github.com/go-vgo/robotgo"
    "github.com/z3oxs/ghost/utils"
)

func Fullscreen(save, format string, clipboard, output, upload bool) {
    screenshot := robotgo.CaptureImg()

    utils.PlaySound("/opt/ghost/screenshot.wav")

    if clipboard {
        utils.SaveToClipboard(screenshot)
    
    }

    if save != "" {
        robotgo.Save(screenshot, save)

    }

    if output {
        utils.OutputToStdout(screenshot, format)

    }

    if upload {
        utils.UploadImage(screenshot, format)

    }
}
