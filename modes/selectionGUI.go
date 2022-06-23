package modes

import (
    "image"
    "sync"

    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/canvas"
    "github.com/z3oxs/ghost/chans"
    "github.com/z3oxs/ghost/utils"
)

func SelectionGUI(save string, clipboard, output, upload, file bool) {
    var selectedDisplay int
    var multipleDisplay bool
    var screenshot image.Image
    var wg sync.WaitGroup
    var screenshots []*canvas.Raster
    var event = make(chan *fyne.KeyEvent)

    app := app.New()

    window := app.NewWindow("Save")
    window.SetMaster()
    window.SetPadded(false)
    window.SetFullScreen(true)

    displays := utils.GetDisplays()

    if len(displays) > 0 {
        multipleDisplay = true
        
        for _, display := range displays {
            displayScreenshot := utils.CaptureRect (
                display.X,
                display.Y,
                display.Width,
                display.Height,
                save,
            )

            image := canvas.NewRasterFromImage(displayScreenshot)

            screenshots = append(screenshots, image)
        }
    }

    if multipleDisplay {
        selectedDisplay = 0

        go chans.KeyEventsMulti (
            event, selectedDisplay,
            window, screenshots, displays,
        )

        window.SetContent(screenshots[selectedDisplay])
    
    } else {
        image := canvas.NewRasterFromImage(utils.CaptureScreen(save))

        go chans.KeyEvents(event)

        window.SetContent(image)
    }

    go func() {
        window.Canvas().SetOnTypedKey(func(e *fyne.KeyEvent) {
            event <- e
        })
    }()

    go func() {
        wg.Add(1)
        defer wg.Done()

        for {
            m := utils.GetMouse()

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

            break
        }

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

        window.Close()
    }()

    window.ShowAndRun()

    wg.Wait()
}
