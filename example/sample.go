package main

import (
	ocr "github.com/cosmo0920/tesseract-ocr-capi-go"
	"fmt"
	"os"
)

const abort = 3

func main() {
	lang := "eng"
	filename := "fixture/golangref.tiff"
	env := ocr.Env()
	version := ocr.Version()
	fmt.Println("tesseract version: " + version)
	api, err := ocr.BaseAPIInit3(env, lang)
	if err != nil {
		os.Exit(abort)
	}
	result := api.BaseAPIProcessPages(filename, nil, 0)
	fmt.Println(result)
}
