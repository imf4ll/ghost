package modes

import (
    "image"

    "github.com/z3oxs/ghost/utils"
)

func Selection(save string, clipboard, output, upload, file bool) {
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

    utils.SaveHandler(clipboard, output, upload, file, screenshot)
}
