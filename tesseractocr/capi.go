package tesseractocr

// #include <stdio.h>
// #include <stdlib.h>
// #include <tesseract/capi.h>
import "C"
import (
	"errors"
	"os"
	"unsafe"
	lept "github.com/cosmo0920/leptonica-capi-go/leptonica"
)

const (
	success = iota
	failure
	uninitialized
	abort
	unknown
)

func Version() string {
	cVersion := C.TessVersion()
	version := C.GoString(cVersion)
	return version
}

func Env() string {
	env := os.Getenv("TESSDATA_PREFIX")
	if env == "" {
		env = "../"
	}
	return env
}

func baseAPICreate() *C.struct_TessBaseAPI {
	return C.TessBaseAPICreate()
}

func BaseAPINew() *TesseractAPI {
	tesseract := baseAPICreate()
	t := &TesseractAPI{api: tesseract}
	return t
}

func (t *TesseractAPI) BaseAPIClear() {
	C.TessBaseAPIClear(t.api)
}

func (t *TesseractAPI) BaseAPIDelete() {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	if !t.disposed {
		C.TessBaseAPIDelete(t.api)
		t.disposed = true
	}
}

func (t *TesseractAPI) BaseAPIInit3(env string, lang string) (C.int, error) {
	cEnv := C.CString(env)
	defer C.free(unsafe.Pointer(cEnv))

	cLang := C.CString(lang)
	defer C.free(unsafe.Pointer(cLang))

	rc := C.TessBaseAPIInit3(t.api, cEnv, cLang)
	if rc != success {
		t.BaseAPIDelete()
		return rc, errors.New("Could not initialize tesseract.")
	}
	return rc, nil
}

func (t *TesseractAPI) BaseAPIProcessPages(filename string, retry_config *C.char, timeout_millisec C.int) string {
	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))

	out := C.TessBaseAPIProcessPages(t.api, cFilename, retry_config, 0)
	result := C.GoString(out)
	return result
}

func (t *TesseractAPI) BaseAPISetVariable(name string, value string) C.int {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	cValue := C.CString(value)
	defer  C.free(unsafe.Pointer(cValue))

	return C.TessBaseAPISetVariable(t.api, cName, cValue)
}

func (t *TesseractAPI) BaseAPISetOutputName(path string) {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))

	C.TessBaseAPISetOutputName(t.api, cPath)
}

func (t *TesseractAPI) BaseAPIPrintVariables() {
	C.TessBaseAPIPrintVariables(t.api, (*C.FILE)(C.stdout))
}

func (t *TesseractAPI) BaseAPIPrintVariablesToFile(name string) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	C.TessBaseAPIPrintVariablesToFile(t.api, cName)
}

func (t *TesseractAPI) BaseAPISetImage(pix *lept.Pix) {
	cPix := pix.RawPix()
	C.TessBaseAPISetImage2(t.api, (*C.struct_Pix)(unsafe.Pointer(cPix)));
}

func (t *TesseractAPI) BaseAPISetSourceResolution(ppi C.int) {
	C.TessBaseAPISetSourceResolution(t.api, ppi);
}

func (t *TesseractAPI) BaseAPISetRectangle(rect Rectangle) {
	C.TessBaseAPISetRectangle(t.api, rect.left, rect.top, rect.width, rect.height)
}