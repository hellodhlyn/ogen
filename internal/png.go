package internal

import (
	"context"
	"io"
	"os"
	"os/exec"
	"strconv"
	"time"
)

const (
	filePath = "/tmp/images"
)

func ConvertSVGToPNG(ctx context.Context, svg io.Reader) (io.Reader, error) {
	makePathIfNotExists(filePath)
	fileName := filePath + "/" + strconv.FormatInt(time.Now().UnixMilli(), 10)
	fileNameSvg := fileName + ".svg"
	fileNamePng := fileName + ".png"

	file, err := os.Create(fileNameSvg)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(file, svg)
	if err != nil {
		return nil, err
	}

	_, err = exec.Command("svgexport", fileNameSvg, fileNamePng).Output()
	if err != nil {
		return nil, err
	}

	pr, pw := io.Pipe()
	go func(_ctx context.Context) {
		defer pw.Close()

		png, _ := os.Open(fileNamePng)
		_, _ = io.Copy(pw, png)
	}(ctx)
	return pr, nil
}

func makePathIfNotExists(pathName string) {
	if _, err := os.Stat(pathName); os.IsNotExist(err) {
		_ = os.MkdirAll(pathName, os.ModePerm)
	}
}
