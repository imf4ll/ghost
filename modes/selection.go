package modes

import (
    "image"

    "github.com/imf4ll/ghost/get"
    "github.com/imf4ll/ghost/save"
    "github.com/imf4ll/ghost/utils"
)

func Selection(filename string, clipboard, output, upload, file bool) {
    var screenshot image.Image 

    m := get.GetMouseAndKeyboard()

    if m.X1 < m.X2 {
        screenshot = utils.CaptureRect (
            m.X1,
            m.Y1,
            m.X2 - m.X1,
            m.Y2 - m.Y1,
            filename,
        )

    } else {
        screenshot = utils.CaptureRect (
            m.X2,
            m.Y2,
            m.X1 - m.X2,
            m.Y1 - m.Y2,
            filename,
        )

    }

    save.SaveHandler(clipboard, output, upload, file, screenshot)
}
