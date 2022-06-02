package utils

import (
    "log"
    "os"
    "net/http"
    "io/ioutil"
    "io"
    "mime/multipart"
    "bytes"
    "image"
    "encoding/json"
    "errors"
    "fmt"

    "github.com/go-vgo/robotgo"
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
    var path string = "/tmp/screenshot.png"

    fmt.Printf("Uploading...\n\n")

    robotgo.Save(screenshot, path)

    contentType, body, err := createFormFile(path)
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

func createFormFile(file string) (string, io.Reader, error) {
    body := new(bytes.Buffer)

    image, err := os.Open(file)
    if err != nil {
        return "", nil, errors.New(err.Error())

    }

    m := multipart.NewWriter(body)
    defer m.Close()

    part, err := m.CreateFormFile("file", file)
    if err != nil {
        return "", nil, errors.New(err.Error())

    }

    _, err = io.Copy(part, image)
    if err != nil {
        return "", nil, errors.New(err.Error())

    }

    return m.FormDataContentType(), body, nil
}
