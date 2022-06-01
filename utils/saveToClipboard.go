package utils

import (
    "image"
    "os/exec"
    "log"

    "github.com/go-vgo/robotgo"
)

func SaveToClipboard(screenshot image.Image) {
    robotgo.Save(screenshot, "/tmp/screenshot.png")

    cmd := exec.Command("xclip", "-selection", "clipboard", "-t", "image/png", "-i", "/tmp/screenshot.png")

    if err := cmd.Run(); err != nil {
        log.Fatal(err)

    }
}
