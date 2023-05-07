package modes

import (
    "log"

    "github.com/imf4ll/ghost/get"
    "github.com/imf4ll/ghost/save"
    "github.com/imf4ll/ghost/utils"
)

func Display(filename string, clipboard, output, upload, file bool, display int) {
    displays := get.GetDisplays()

    if display > len(displays) {
        log.Fatal("Invalid screen selected.")

    }

    displaySelected := displays[display]

    screenshot := utils.CaptureRect (
        displaySelected.X,
        displaySelected.Y,
        displaySelected.Width,
        displaySelected.Height,
        filename,
    )

    save.SaveHandler(clipboard, output, upload, file, screenshot)
}
