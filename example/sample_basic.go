package main

import (
	ocr "../tesseractocr"
	lept "github.com/cosmo0920/leptonica-capi-go/leptonica"
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
	pix, _ := lept.PixRead(filename)
	api.BaseAPISetImage(pix)
	text := api.BaseAPIGetUTF8Text()
	fmt.Println(text)
}
