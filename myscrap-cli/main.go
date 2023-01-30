package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

// var format string

// func init() {
// 	flag.StringVar(&format, "f", "", "変換する画像フォーマットを指定してください")
// }

func main() {

	args := os.Args
	cmd := args[1]
	param := args[2]

	switch cmd {
	case "clip":
		clipper := NewClipper()
		if err := clipper.ClipURL(param); err != nil {
			log.Fatalln(err)
		}
	}
}

type ClipperInterface interface {
	ClipURL(param string) error
}

type Clipper struct{}

func NewClipper() ClipperInterface {
	return &Clipper{}
}

const (
	ACCESS_KEY_ID     = "0f73919d-4d81-4c79-b982-e8d7a6f2a3ce"
	SECRET_ACCESS_KEY = "c131ef5c-70b9-4051-bce0-f9ffec7705e8"
	ENDPOINT          = "http://localhost:8080"
	CLIP_COMMAND      = "clip"
)

func (u *Clipper) ClipURL(param string) error {

	url, err := url.JoinPath(ENDPOINT, CLIP_COMMAND)
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(
		&ClipRequestParam{
			AccessKeyID:     ACCESS_KEY_ID,
			SecretAccessKey: SECRET_ACCESS_KEY,
			URL:             param,
		},
	); err != nil {
		return err
	}

	resp, err := http.Post(url, http.DetectContentType(buf.Bytes()), &buf)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println(resp.StatusCode)
	fmt.Println(string(body))

	return nil
}

type ClipRequestParam struct {
	AccessKeyID     string
	SecretAccessKey string
	URL             string
}
