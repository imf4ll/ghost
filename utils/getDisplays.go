package utils

import (
    "os/exec"
    "strconv"
    "strings"
    "log"
)

type Display struct {
    Name string
    Width int
    Height int
    X int
    Y int
}

func GetDisplays() []Display {
    var displays []Display
    var displaysSizes []string

    output, err := exec.Command("bash", "-c", "xrandr | grep ' connected'").Output()
    if err != nil {
        log.Fatal(err)

    }

    displaysInformation := strings.Split(string(output), "\n")
    displaysInformation = displaysInformation[:len(displaysInformation) - 1]

    for _, i := range displaysInformation {
        displayInfo := strings.Split(i, " ")

        if displayInfo[2] != "primary" {
            displaysSizes = strings.Split(displayInfo[2], "+")
        
        } else {
            displaysSizes = strings.Split(displayInfo[3], "+")

        }

        monitorResolution := strings.Split(displaysSizes[0], "x")

        width, _ := strconv.Atoi(monitorResolution[0])
        height, _ := strconv.Atoi(monitorResolution[1])
        x, _ := strconv.Atoi(displaysSizes[1])
        y, _ := strconv.Atoi(displaysSizes[2])

        displays = append(displays, Display {
            Name: displayInfo[0],
            Width: width,
            Height: height,
            X: x,
            Y: y,
        })
    }

    return displays
}
