package modes

import "github.com/z3oxs/ghost/utils"

func Fullscreen(save string, clipboard, output, upload, file bool) {
    screenshot := utils.CaptureScreen(save)

    if clipboard {
        utils.SaveToClipboard(screenshot)
    
    }

    if output {
        utils.OutputToStdout(screenshot)

    }

    if upload {
        utils.UploadImage(screenshot)

    }

    if file {
        utils.SaveImage(screenshot)

    }
}

