package tesseractocr

// #include <stdio.h>
// #include <stdlib.h>
// #include <tesseract/capi.h>
import "C"
import (
	"errors"
	lept "github.com/cosmo0920/leptonica-capi-go"
	"os"
	"runtime"
	"unsafe"
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

func (t *TesseractAPI) BaseAPIClear() {
	C.TessBaseAPIClear(t.api)
}

func baseAPINew() *TesseractAPI {
	tesseract := C.TessBaseAPICreate()
	t := &TesseractAPI{api: tesseract}
	return t
}

func (t *TesseractAPI) baseAPIDelete() {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	if !t.disposed {
		C.TessBaseAPIDelete(t.api)
		t.disposed = true
	}
}

func (t *TesseractAPI) baseAPIEnd() {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	if !t.disposed {
		C.TessBaseAPIEnd(t.api)
		t.disposed = true
	}
}

func (t *TesseractAPI) finalize() {
	t.baseAPIEnd()
	t.baseAPIDelete()
}

func BaseAPIInit2(env string, lang string, oem TessOcrEngineMode) (*TesseractAPI, error) {
	cEnv := C.CString(env)
	defer C.free(unsafe.Pointer(cEnv))

	cLang := C.CString(lang)
	defer C.free(unsafe.Pointer(cLang))

	instance := baseAPINew()

	rc := C.TessBaseAPIInit2(instance.api, cEnv, cLang, C.TessOcrEngineMode(oem))
	if rc != success {
		return nil, errors.New("Could not initialize tesseract.")
	}

	runtime.SetFinalizer(instance, (*TesseractAPI).finalize)
	return instance, nil
}

func BaseAPIInit3(env string, lang string) (*TesseractAPI, error) {
	cEnv := C.CString(env)
	defer C.free(unsafe.Pointer(cEnv))

	cLang := C.CString(lang)
	defer C.free(unsafe.Pointer(cLang))

	instance := baseAPINew()

	rc := C.TessBaseAPIInit3(instance.api, cEnv, cLang)
	if rc != success {
		return nil, errors.New("Could not initialize tesseract.")
	}

	runtime.SetFinalizer(instance, (*TesseractAPI).finalize)
	return instance, nil
}

func (t *TesseractAPI) BaseAPIProcessPages(filename string, retry_config *C.char, timeout_millisec int) string {
	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))

	out := C.TessBaseAPIProcessPages(t.api, cFilename, retry_config,
		C.int(timeout_millisec))
	result := C.GoString(out)
	return result
}

func (t *TesseractAPI) BaseAPIProcessPage(filename string, pix *lept.Pix, page_index int, retry_config *C.char, timeout_millisec int) string {
	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))
	cPix := pix.RawPix()

	out := C.TessBaseAPIProcessPage(
		t.api, (*C.struct_Pix)(unsafe.Pointer(cPix)),
		C.int(page_index), cFilename,
		retry_config, C.int(timeout_millisec))
	result := C.GoString(out)
	return result
}

func (t *TesseractAPI) BaseAPISetVariable(name string, value string) C.int {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	cValue := C.CString(value)
	defer C.free(unsafe.Pointer(cValue))

	return C.TessBaseAPISetVariable(t.api, cName, cValue)
}

func (t *TesseractAPI) BaseAPIGetIntVariable(name string) C.int {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	var cValue C.int
	C.TessBaseAPIGetIntVariable(t.api, cName, &cValue)

	return cValue
}

func (t *TesseractAPI) BaseAPIGetBoolVariable(name string) C.int {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	var cValue C.int
	C.TessBaseAPIGetBoolVariable(t.api, cName, &cValue)

	return cValue
}

func (t *TesseractAPI) BaseAPIGetDoubleVariable(name string) C.double {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	var cValue C.double
	C.TessBaseAPIGetDoubleVariable(t.api, cName, &cValue)

	return cValue
}

func (t *TesseractAPI) BaseAPIGetStringVariable(name string) (string, error) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	cValue := C.TessBaseAPIGetStringVariable(t.api, cName)

	goValue := C.GoString(cValue)
	if goValue == "" {
		return "", errors.New("Could not get variable contents")
	}

	return goValue, nil
}

func (t *TesseractAPI) BaseAPISetInputName(name string) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	C.TessBaseAPISetInputName(t.api, cName)
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
	C.TessBaseAPISetImage2(t.api, (*C.struct_Pix)(unsafe.Pointer(cPix)))
}

func (t *TesseractAPI) BaseAPIGetUTF8Text() string {
	cText := C.TessBaseAPIGetUTF8Text(t.api)
	defer C.free(unsafe.Pointer(cText))

	return C.GoString(cText)
}

// Since 3.03? In 3.02, this feature is not enabled....
// func (t *TesseractAPI) BaseAPIGetOpenCLDevice() (**C.void) {
// 	var device (**C.void)
// 	C.TessBaseAPIGetOpenCLDevice(t.api, device)
// 	return device
// }

func (t *TesseractAPI) BaseAPISetSourceResolution(ppi int) {
	C.TessBaseAPISetSourceResolution(t.api, C.int(ppi))
}

func (t *TesseractAPI) BaseAPISetRectangle(rect Rectangle) {
	C.TessBaseAPISetRectangle(t.api,
		C.int(rect.left), C.int(rect.top),
		C.int(rect.width), C.int(rect.height))
}

func (t *TesseractAPI) BaseAPISetSegMode(mode TessPageSegMode) {
	C.TessBaseAPISetPageSegMode(t.api, C.TessPageSegMode(mode))
}

func (t *TesseractAPI) BaseAPIGetSegMode() TessPageSegMode {
	seg := C.TessBaseAPIGetPageSegMode(t.api)
	return TessPageSegMode(seg)
}

func (t *TesseractAPI) BaseAPIGetComponentImages(level TessPageIteratorLevel,
	flag int, pixa **C.struct_Pixa, blockids **C.int) *C.BOXA {
	boxaImg := C.TessBaseAPIGetComponentImages(t.api, C.TessPageIteratorLevel(level), C.int(flag), pixa, blockids)
	return boxaImg
}
