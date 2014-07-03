package tesseractocr

// #include <stdio.h>
// #include <tesseract/capi.h>
import "C"
import (
	"errors"
	"os"
)

const (
	success = iota
	failure
	uninitialized
	abort
	unknown
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

func BaseAPICreate() *C.struct_TessBaseAPI {
	return C.TessBaseAPICreate()
}

func BaseAPINew() *TesseractAPI {
	tesseract := BaseAPICreate()
	t := &TesseractAPI{tesseract}
	return t
}

func (t *TesseractAPI) BaseAPIDelete() error {
	C.TessBaseAPIDelete(t.api)
	return nil
}

func (t *TesseractAPI) BaseAPIInit3(env string, lang string) (C.int, error) {
	cenv := C.CString(env)
	clang := C.CString(lang)
	rc := C.TessBaseAPIInit3(t.api, cenv, clang)
	if rc != success {
		_ = t.BaseAPIDelete()
		return rc, errors.New("Could not initialize tesseract.")
	}
	return rc, nil
}

func (t *TesseractAPI) BaseAPIProcessPages(filename string, retry_config *C.char, timeout_millisec C.int) string {
	cfilename := C.CString(filename)
	out := C.TessBaseAPIProcessPages(t.api, cfilename, retry_config, 0)
	result := C.GoString(out)
	return result
}

func (t *TesseractAPI) BaseAPISetVariable(name string, value string) C.int {
	cname := C.CString(name)
	cvalue := C.CString(value)
	return C.TessBaseAPISetVariable(t.api, cname, cvalue)
}

func (t *TesseractAPI) BaseAPISetOutputName(path string) {
	cpath := C.CString(path)
	C.TessBaseAPISetOutputName(t.api, cpath)
}

func (t *TesseractAPI) BaseAPIPrintVariablesToFile(name string) {
	cname := C.CString(name)
	C.TessBaseAPIPrintVariablesToFile(t.api, cname)
}

func (t *TesseractAPI) BaseAPISetImage2(pix *C.struct_Pix) {
	C.TessBaseAPISetImage2(t.api, pix);
}

func (t *TesseractAPI) BaseAPISetSourceResolution(ppi C.int) {
	C. TessBaseAPISetSourceResolution(t.api, ppi);
}