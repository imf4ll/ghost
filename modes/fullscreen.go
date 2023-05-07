package modes

import (
    "github.com/imf4ll/ghost/save"
    "github.com/imf4ll/ghost/utils"
)

func Fullscreen(filename string, clipboard, output, upload, file bool) {
    screenshot := utils.CaptureScreen(filename)

    save.SaveHandler(clipboard, output, upload, file, screenshot)
}

