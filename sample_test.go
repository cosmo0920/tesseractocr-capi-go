package tesseractocr_test

import (
	ocr "./tesseractocr"
	lept "github.com/cosmo0920/leptonica-capi-go/leptonica"
	"fmt"
	"os"
)

const abort = 3

func setupExampleTesseractAPI() (*ocr.TesseractAPI) {
	env := ocr.Env()
	lang := "eng"
	api, err := ocr.BaseAPIInit3(env, lang)
	if err != nil {
		os.Exit(abort)
	}
	return api
}

func ExampleUsage() {
	filename := "fixture/golangref.tiff"
	api := setupExampleTesseractAPI()
	result := api.BaseAPIProcessPages(filename, nil, 0)
	fmt.Println(result)
	// output: Function pointer callbacks
	//
	// C code can call exported Go functions with their explicit name. But if a C—program wants a function pointer, a gateway function has to be written.
	// This is because we can't take the address of a Go function and give that to C—code since the cgo tool will generate a stub in C that should be
	// called. The following example shows how to integrate with C code wanting a function pointer of a give type.
	//
	// Place these source ﬁles under $GOPATH/src/ccallbacks/. Compile and run with:
}

func ExampleBasicUsage() {
	filename := "fixture/golangref.tiff"
	api := setupExampleTesseractAPI()
	pix, _ := lept.PixRead(filename)
	api.BaseAPISetImage(pix)
	text := api.BaseAPIGetUTF8Text()
	fmt.Println(text)
	// output: Function pointer callbacks
	//
	// C code can call exported Go functions with their explicit name. But if a C—program wants a function pointer, a gateway function has to be written.
	// This is because we can't take the address of a Go function and give that to C—code since the cgo tool will generate a stub in C that should be
	// called. The following example shows how to integrate with C code wanting a function pointer of a give type.
	//
	// Place these source ﬁles under $GOPATH/src/ccallbacks/. Compile and run with:
}
