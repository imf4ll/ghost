package utils

import (
    "fmt"
    "image"
    "image/png"
    "log"
    "os"
    "time"
)

func SaveImage(screenshot image.Image) {
    var save, hour, minute, second, month string

    now := time.Now()
    month = fmt.Sprintf("%v", int(now.Month()))
    hour = fmt.Sprintf("%v", now.Hour())
    minute = fmt.Sprintf("%v", now.Minute())
    second = fmt.Sprintf("%v", now.Second())

    if int(now.Month()) < 10 {
        month = fmt.Sprintf("0%v", int(now.Month()))

    }

    if now.Hour() < 10 {
        hour = fmt.Sprintf("0%v", now.Hour())

    }

    if now.Minute() < 10 {
        minute = fmt.Sprintf("0%v", now.Minute())

    }

    if now.Second() < 10 {
        second = fmt.Sprintf("0%v", now.Second())

    }

    save = fmt.Sprintf("%v/Images/%v%v%v_%v%v%v.png",
        os.Getenv("HOME"),
        now.Year(),
        month,
        now.Day(),
        hour,
        minute,
        second,
    )

    file, err := os.Create(save)
    if err != nil {
        log.Fatal(err)

    }

    err = png.Encode(file, screenshot)
    if err != nil {
        log.Fatal(err)

    }
}
