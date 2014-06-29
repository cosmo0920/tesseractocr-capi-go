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

func BaseAPICreate() {
	api = C.TessBaseAPICreate()
}

func BaseAPIDelete() error {
	if api == nil {
		return errors.New("call BaseAPICreate() first")
	}
	C.TessBaseAPIDelete(api)
	return nil
}

func BaseAPIInit3(env string, lang string) (C.int, error) {
	cenv := C.CString(env)
	clang := C.CString(lang)
	if api == nil {
		return uninitialized, errors.New("call BaseAPICreate() first")
	}
	rc := C.TessBaseAPIInit3(api, cenv, clang)
	if rc != success {
		_ = BaseAPIDelete()
		return rc, errors.New("Could not initialize tesseract.")
	}
	return rc, nil
}

func BaseAPIProcessPages(filename string, retry_config *C.char, timeout_millisec C.int) (string, error) {
	cfilename := C.CString(filename)
	if api == nil {
		return "", errors.New("call BaseAPICreate() first")
	}
	out := C.TessBaseAPIProcessPages(api, cfilename, retry_config, 0)
	result := C.GoString(out)
	return result, nil
}

func BaseAPISetVariable(name string, value string) (C.int, error) {
	cname := C.CString(name)
	cvalue := C.CString(value)
	if api == nil {
		return uninitialized, errors.New("call BaseAPICreate() first")
	}
	return C.TessBaseAPISetVariable(api, cname, cvalue), nil
}

func BaseAPISetOutputName(path string) (error) {
	cpath := C.CString(path)
	if api == nil {
		return errors.New("call BaseAPICreate() first")
	}
	C.TessBaseAPISetOutputName(api, cpath)
	return nil
}

func BaseAPIPrintVariablesToFile(name string) (error) {
	cname := C.CString(name)
	if api == nil {
		return errors.New("call BaseAPICreate() first")
	}
	C.TessBaseAPIPrintVariablesToFile(api, cname)
	return nil
}
