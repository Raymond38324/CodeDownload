package main

import (
	"archive/zip"
	"fmt"
	"io"
	"strings"
)

func ExtraFile(path string) (string, error) {
	reader, err := zip.OpenReader(path)
	var content strings.Builder
	if err != nil {
		return "", err
	}
	defer reader.Close()
	for _, file := range reader.File {
		rc, err := file.Open()
		if err != nil {
			return "", err
		}
		if strings.Contains(file.Name, "2018"){
			io.Copy(&content, rc)
			return content.String(), nil
		}else {
			continue
		}
	}
	return "", fmt.Errorf("没有符合要求的文件")
}