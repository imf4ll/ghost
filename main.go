package main

import (
    "flag"
    "os"
    "fmt"
    "log"

    "github.com/z3oxs/ghost/modes"
)

var (
    fullscreen,
    selection,
    selectiongui,
    clipboard,
    output,
    upload,
    version,
    file bool
    display int
    save string
)

func init() {
    flag.BoolVar(&fullscreen, "f", false, "Fullscreen mode")
    flag.BoolVar(&selection, "s", false, "Selection mode")
    flag.BoolVar(&selectiongui, "g", false, "Selection with GUI mode")
    flag.IntVar(&display, "d", -1, "Display mode")
    flag.BoolVar(&clipboard, "c", false, "Save to clipboard")
    flag.StringVar(&save, "S", "/tmp/screenshot.png", "Save to file")
    flag.BoolVar(&file, "F", false, "Save to file on default images path")
    flag.BoolVar(&output, "o", false, "Output to stdout")
    flag.BoolVar(&upload, "u", false, "Submit to AnonFiles")
    flag.BoolVar(&version, "v", false, "Version")
    
    flag.Usage = func() {
        fmt.Fprintf(os.Stdout, `Usage:
    ghost [flags]

Description:
    Simple tool to take screenshots

Flags:
    -f    Fullscreen mode
    -s    Selection mode
    -g    Selection with GUI mode
    -d    Display mode (int)
    -c    Save to clipboard
    -S    Save to file (string)
    -F    Save to file on default images path
    -o    Output to stdout
    -u    Submit to AnonFiles
    -v    Print version
`)
    }

    flag.Parse()
}

func main() {
    if version {
        fmt.Println("\nCurrent version:\n    v1.1.0")

        os.Exit(3)
    }

    if (!clipboard && !file && !output && !upload) {
        log.Fatalln("You need to specify at least a save method, check 'ghost --help'.")

    }

    if fullscreen {
       modes.Fullscreen(save, clipboard, output, upload, file)

    } else if selection {
        modes.Selection(save, clipboard, output, upload, file)

    } else if selectiongui {
        modes.SelectionGUI(save, clipboard, output, upload, file)

    } else if display != -1 {
        modes.Display(save, clipboard, output, upload, file, display)

    } else {
        flag.Usage()

    }
}
