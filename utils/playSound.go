package utils

import (
    "os"
    "log"
    "fmt"
    "time"
    "runtime"

    "github.com/faiface/beep"
    "github.com/faiface/beep/wav"
    "github.com/faiface/beep/speaker"
)

func PlaySound(sound string) {
    var file * os.File

    done := make(chan bool)

    if runtime.GOOS == "windows" {
        file, _ = os.Open(fmt.Sprintf("%s\\AppData\\Roaming\\ghost\\%s", os.Getenv("HOME"), sound))
    
    } else if runtime.GOOS == "linux" {
        file, _ = os.Open(fmt.Sprintf("/opt/ghost/%s", sound))

    }

    streamer, format, err := wav.Decode(file)
    if err != nil {
        log.Fatal(err)

    }

    defer streamer.Close()

    speaker.Init(format.SampleRate, format.SampleRate.N(time.Second / 10))

    speaker.Play(beep.Seq(streamer, beep.Callback(func() {
        done <- true

    })))

    <- done
}
