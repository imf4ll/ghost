package utils

import (
    "os"
    "fmt"
    "runtime"
)

func CheckSave(save string) bool {
    switch runtime.GOOS {
        case "linux":
            if save == "/tmp/screenshot.png" {
                return true

            }

        case "windows":
            if save == fmt.Sprintf("%s\\AppData\\Local\\Temp\\screenshot.png", os.Getenv("HOME")) {
                return true

            }
    }

    return false
}
