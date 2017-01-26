package main

import (
	ocr "github.com/cosmo0920/tesseractocr-capi-go"
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
	oem := ocr.OEM_CUBE_ONLY
	api, err := ocr.BaseAPIInit2(env, lang, oem)
	if err != nil {
		os.Exit(abort)
	}
	result := api.BaseAPIProcessPages(filename, nil, 0)
	fmt.Println(result)
}
