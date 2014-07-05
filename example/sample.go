package main

import (
	ocr "../tesseractocr"
	"fmt"
	"os"
)

const abort = 3

func main() {
	lang := "eng"
	filename := "fixture/golangref.png"
	env := ocr.Env()
	version := ocr.Version()
	fmt.Println("tesseract version: " + version)
	api := ocr.BaseAPINew()
	_, err := api.BaseAPIInit3(env, lang)
	if err != nil {
		os.Exit(abort)
	}
	result := api.BaseAPIProcessPages(filename, nil, 0)
	fmt.Println(result)
}
