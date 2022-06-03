package chans

import (
    "os"

    "fyne.io/fyne/v2"
)

func KeyEvents (event chan *fyne.KeyEvent) {
    for {
        select {
            case e := <- event:
                if e.Name == fyne.KeyEscape {
                    os.Exit(3)
                }
        }
    }
}
