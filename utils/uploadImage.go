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

func UploadImage(screenshot image.Image) {
    var response Response
    var image = new(bytes.Buffer)

    fmt.Printf("Uploading...\n\n")

    err := png.Encode(image, screenshot)
    if err != nil {
        log.Fatal(err)

    }

    contentType, body, err := createForm(image)
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

func createForm(image io.Reader) (string, io.Reader, error) {
    body := new(bytes.Buffer)

    m := multipart.NewWriter(body)
    defer m.Close()

    part, err := m.CreateFormFile("file", "screenshot.png")
    if err != nil {
        return "", nil, errors.New(err.Error())

    }

    _, err = io.Copy(part, image)
    if err != nil {
        return "", nil, errors.New(err.Error())

    }

    return m.FormDataContentType(), body, nil
}
