package main

import (
    "flag"
    "runtime"
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
    version bool
    display int
    save string
)

func init() {
    var saveString string

    switch runtime.GOOS {
        case "linux": saveString = "/tmp/screenshot.png"

        case "windows":
            if os.Getenv("USERPROFILE") != "" {
                saveString = fmt.Sprintf("%s\\AppData\\Local\\Temp\\screenshot.png", os.Getenv("USERPROFILE"))
            
            } else {
                saveString = fmt.Sprintf("%s\\AppData\\Local\\Temp\\screenshot.png", os.Getenv("HOME"))

            }

        case "darwin": panic("Not yet implemented.")
    } 

    flag.BoolVar(&fullscreen, "f", false, "Fullscreen mode")
    flag.BoolVar(&selection, "s", false, "Selection mode")
    flag.BoolVar(&selectiongui, "g", false, "Selection with GUI mode")
    flag.IntVar(&display, "d", -1, "Display mode")
    flag.BoolVar(&clipboard, "c", false, "Save to clipboard")
    flag.StringVar(&save, "S", saveString, "Save to file")
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
    -S    Save to file (default: %s) (string)
    -o    Output to stdout
    -u    Submit to AnonFiles
    -v    Print version
`, saveString)
    }

    flag.Parse()
}

func main() {
    if version {
        fmt.Println("\nCurrent version:\n    v1.0.9")

        os.Exit(3)
    }

    if (!clipboard && save == "" && !output && !upload) {
        log.Fatalln("You need to specify at least a save method, check 'ghost --help'.")

    }

    if fullscreen {
       modes.Fullscreen(save, clipboard, output, upload)

    } else if selection {
        modes.Selection(save, clipboard, output, upload)

    } else if selectiongui {
        modes.SelectionGUI(save, clipboard, output, upload)

    } else if display != -1 {
        modes.Display(save, clipboard, output, upload, display)

    } else {
        flag.Usage()

    }
}
