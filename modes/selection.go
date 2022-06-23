package modes

import (
    "image"

    "github.com/z3oxs/ghost/utils"
)

func Selection(save string, clipboard, output, upload bool) {
    var screenshot image.Image 

    m := utils.GetMouseAndKeyboard()

    if m.X1 < m.X2 {
        screenshot = utils.CaptureRect (
            m.X1,
            m.Y1,
            m.X2 - m.X1,
            m.Y2 - m.Y1,
            save,
        )

    } else {
        screenshot = utils.CaptureRect (
            m.X2,
            m.Y2,
            m.X1 - m.X2,
            m.Y1 - m.Y2,
            save,
        )

    }

    utils.SaveImage(screenshot, save)

    if clipboard {
        utils.SaveToClipboard(screenshot)

    }

    if output {
        utils.OutputToStdout(screenshot)

    }

    if upload {
        utils.UploadImage(screenshot)

    }
}
