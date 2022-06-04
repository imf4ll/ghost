package utils

import (
    "log"
    "net/http"
    "io/ioutil"
    "io"
    "mime/multipart"
    "bytes"
    "image"
    "image/png"
    "image/jpeg"
    "encoding/json"
    "errors"
    "fmt"
)

type Response struct {
    Status bool `json:"status"`
    Data struct {
        File struct {
            URL struct {
                Full string `json:"full"`
                Short string `json:"short"`
            } `json:"url"`
            Metadata struct {
                ID string `json:"id"`
                Name string `json:"name"`
                Size struct {
                    Bytes int `json:"bytes"`
                    Readable string `json:"readable"`
                } `json:"size"`
            } `json:"metadata"`
        } `json:"file"`
    } `json:"data"`
}

func UploadImage(screenshot image.Image, format string) {
    var response Response
    var image = new(bytes.Buffer)

    fmt.Printf("Uploading...\n\n")

    switch format {
        case "png":
            err := png.Encode(image, screenshot)
            if err != nil {
                log.Fatal(err)

            }

        case "jpg":
            err := jpeg.Encode(image, screenshot, &jpeg.Options {
                Quality: 100,
            })
            if err != nil {
                log.Fatal(err)

            }
    }

    contentType, body, err := createForm(image, format)
    if err != nil {
        log.Fatal(err)

    }

    resp, err := http.Post("https://api.anonfiles.com/upload", contentType, body)
    if err != nil {
        log.Fatal(err)

    }

    data, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)

    }

    json.Unmarshal(data, &response)

    fmt.Printf("ID: %s\nSize: %s\nURL: %s\n",
        response.Data.File.Metadata.ID,
        response.Data.File.Metadata.Size.Readable,
        response.Data.File.URL.Full,
    )
}

func createForm(image io.Reader, format string) (string, io.Reader, error) {
    body := new(bytes.Buffer)

    m := multipart.NewWriter(body)
    defer m.Close()

    part, err := m.CreateFormFile("file", fmt.Sprintf("screenshot.%s", format))
    if err != nil {
        return "", nil, errors.New(err.Error())

    }

    _, err = io.Copy(part, image)
    if err != nil {
        return "", nil, errors.New(err.Error())

    }

    return m.FormDataContentType(), body, nil
}
