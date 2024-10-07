package main

import (
	"bytes"
	"fmt"
	"io"
	"os"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func extractThumbnail(inputPath string, seconds float64) (io.Reader, error) {
	var buf bytes.Buffer
	err := ffmpeg.Input(inputPath, ffmpeg.KwArgs{"ss": seconds}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(&buf, os.Stdout).
		Run()
	if err != nil {
		return nil, err
	}
	return &buf, nil
}

func main() {
	inputPath := "input/sample.mp4"

	reader, err := extractThumbnail(inputPath, 3.5)
	if err != nil {
        fmt.Println("Failed to extract thumbnail")
        return
    }

	outputPath := "output/thumbnail.jpg"
    file, err := os.Create(outputPath)
    if err != nil {
        fmt.Printf("Error creating file: %v\n", err)
        return
    }
    defer file.Close()

    _, err = io.Copy(file, reader)
    if err != nil {
        fmt.Printf("Error saving thumbnail: %v\n", err)
        return
    }

	fmt.Println("success to extract image")
}
