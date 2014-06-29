package tesseractocr

// #include <stdio.h>
// #include <tesseract/capi.h>
import "C"
import (
	"errors"
	"os"
)

var (
	api *C.struct_TessBaseAPI
)

func Version() string {
	c_version := C.TessVersion()
	version := C.GoString(c_version)
	return version
}

func Env() string {
	env := os.Getenv("TESSDATA_PREFIX")
	if env == "" {
		env = "../"
	}
	return env
}

func BaseAPICreate() {
	api = C.TessBaseAPICreate()
}

func BaseAPIDelete() {
	C.TessBaseAPIDelete(api)
}

func BaseAPIInit3(env string, lang string) (C.int, error) {
	cenv := C.CString(env)
	clang := C.CString(lang)
	rc := C.TessBaseAPIInit3(api, cenv, clang)
	if rc != 0 {
		BaseAPIDelete()
		return rc, errors.New("Could not initialize tesseract.")
	}
	return rc, nil
}

func BaseAPIProcessPages(filename string, retry_config *C.char, timeout_millisec C.int) string {
	cfilename := C.CString(filename)
	out := C.TessBaseAPIProcessPages(api, cfilename, retry_config, 0)
	result := C.GoString(out)
	return result
}
