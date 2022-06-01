package utils

import (
    "os"
    "log"
    "time"

    "github.com/faiface/beep"
    "github.com/faiface/beep/wav"
    "github.com/faiface/beep/speaker"
)

func PlaySound(sound string) {
    done := make(chan bool)

    file, err := os.Open(sound)
    if err != nil {
        log.Fatal(err)

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
