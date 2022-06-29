package modes

import "github.com/z3oxs/ghost/utils"

func Fullscreen(save string, clipboard, output, upload, file bool) {
    screenshot := utils.CaptureScreen(save)

    utils.SaveHandler(clipboard, output, upload, file, screenshot)
}

