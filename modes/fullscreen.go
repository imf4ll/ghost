package modes

import (
    "github.com/go-vgo/robotgo"
    "github.com/z3oxs/goshot/utils"
)

func Fullscreen(save string, clipboard, output bool) {
    screenshot := robotgo.CaptureImg()

    if clipboard {
        utils.SaveToClipboard(screenshot)
    
    } else if save != "" {
        robotgo.Save(screenshot, save)

    } else {
        utils.SaveToClipboard(screenshot)

    }

    if output {
        utils.OutputToStdout(screenshot)
    }
    
    utils.PlaySound("/opt/goshot/screenshot.wav")
}
