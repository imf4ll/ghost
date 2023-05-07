package save

import (
    "log"
    "image"
    "os/exec"
    "strings"
)

func SaveToClipboard(screenshot image.Image) {
    whereIsXClip, err := exec.Command("which", "xclip").Output()
    if err != nil {
        log.Fatal(err)

    }
    
    if strings.Contains(string(whereIsXClip), "not found") {
        log.Fatal("XClip was not found, are you sure that you have installed?")

    }

    cmd := exec.Command("xclip", "-sel", "clipboard", "-t", "image/png", "-i", "/tmp/screenshot.png")

    if err := cmd.Run(); err != nil {
        log.Fatal(err)

    }
}
