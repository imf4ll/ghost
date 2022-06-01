package modes

import (
    "image"
    "log"

    "github.com/z3oxs/goshot/utils"
    "github.com/go-vgo/robotgo"
)

func Monitor(save string, clipboard bool, monitor int, output bool) {
    var screenshot image.Image

    monitors := utils.GetDisplays()

    if monitor > len(monitors) {
        log.Fatal("Invalid screen selected.")

    }

    monitorSelected := monitors[monitor]

    screenshot = robotgo.CaptureImg(
        monitorSelected.X,
        monitorSelected.Y,
        monitorSelected.Width,
        monitorSelected.Height,
    )

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
