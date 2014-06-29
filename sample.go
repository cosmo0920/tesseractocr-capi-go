package main

import (
	ocr "./tesseractocr"
	"fmt"
	"os"
)

func main() {
	lang := "eng"
	filename := "<add your png/jpeg image>"
	env := ocr.Env()
	version := ocr.Version()
	fmt.Println("tesseract version: " + version)
	ocr.BaseAPICreate()
	_, err := ocr.BaseAPIInit3(env, lang)
	if err != nil {
		os.Exit(3)
	}
	result, err := ocr.BaseAPIProcessPages(filename, nil, 0)
	fmt.Println(result)
}
