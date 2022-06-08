package modes

import (
    "sync"
    "image"

    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/canvas"
    "github.com/go-vgo/robotgo"
    "github.com/z3oxs/ghost/chans"
    "github.com/z3oxs/ghost/utils"
)

func SelectionGUI(save string, clipboard, output, upload bool) {
    var x1, x2, y1, y2 int
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
            displayScreenshot := robotgo.CaptureImg (
                display.X,
                display.Y,
                display.Width,
                display.Height,
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
        image := canvas.NewRasterFromImage(robotgo.CaptureImg())

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
            mouseFirstClick := robotgo.AddMouse("left")

            if mouseFirstClick {
                x1, y1 = robotgo.GetMousePos()

            }

            mouseLastClick := robotgo.AddMouse("left")

            if mouseLastClick {
                x2, y2 = robotgo.GetMousePos()

            } 

            if x1 < x2 {
                screenshot = utils.CaptureRect (
                    x1,
                    y1,
                    x2 - x1,
                    y2 - y1,
                    save,
                )

            } else {
                screenshot = utils.CaptureRect (
                    x2,
                    y2,
                    x1 - x2,
                    y1 - y2,
                    save,
                )

            }

            break
        }

        utils.PlaySound("screenshot.wav")

        if utils.CheckSave(save) {
           utils.SaveImage(screenshot, save) 

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

        window.Close()
    }()

    window.ShowAndRun()

    wg.Wait()
}
