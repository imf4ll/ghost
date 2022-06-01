package modes

import (
    "image"
    "log"

    "github.com/z3oxs/goshot/utils"

    "github.com/go-vgo/robotgo"
)

func Selection(save string, clipboard, output bool) {
    var x1, x2, y1, y2 int
    var screenshot image.Image

    mouseFirstClick := robotgo.AddMouse("left")

    if mouseFirstClick {
        x1, y1 = robotgo.GetMousePos()

    }

    mouseLastClick := robotgo.AddMouse("left")

    if mouseLastClick {
        x2, y2 = robotgo.GetMousePos()

    }

    if x1 < x2 {
        screenshot = robotgo.CaptureImg(x1, y1, x2 - x1, y2 - y1) 

    } else {
        screenshot = robotgo.CaptureImg(x2, y2, x1 - x2, y1 - y2)

    }

    if clipboard {
        utils.SaveToClipboard(screenshot)
    
    } else if save != "" {
        if err := robotgo.Save(screenshot, save); err != nil {
            log.Fatal(err)

        }
    
    } else {
        utils.SaveToClipboard(screenshot)

    }

    if output {
        utils.OutputToStdout(screenshot)
    }

    utils.PlaySound("/opt/goshot/screenshot.wav")
}
