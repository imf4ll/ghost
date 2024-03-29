package chans

import (
    "os"

    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/canvas"
    "github.com/imf4ll/ghost/types"
)

func KeyEventsMulti (
    event chan *fyne.KeyEvent, selectedDisplay int,
    window fyne.Window, screenshots []*canvas.Raster, displays []types.Display,
) {
    for {
        select {
            case e := <- event:
                if e.Name == fyne.KeyLeft && selectedDisplay > 0 {
                    selectedDisplay -= 1

                    window.SetContent(screenshots[selectedDisplay])

                } else if e.Name == fyne.KeyRight && selectedDisplay < len(displays) - 1 {
                    selectedDisplay += 1

                    window.SetContent(screenshots[selectedDisplay])
                
                } else if e.Name == fyne.KeyEscape {
                    os.Exit(3)

                }
        }
    }
}
