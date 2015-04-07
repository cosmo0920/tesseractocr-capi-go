package main

import (
	ocr "github.com/cosmo0920/tesseractocr-capi-go"
	lept "github.com/cosmo0920/leptonica-capi-go"
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
	api, err := ocr.BaseAPIInit3(env, lang)
	if err != nil {
		os.Exit(abort)
	}
	pix, err := lept.PixRead(filename)
	if err != nil {
		os.Exit(abort)
	}
	api.BaseAPISetImage(pix)
	text := api.BaseAPIGetUTF8Text()
	fmt.Println(text)
}
