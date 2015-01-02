package tesseractocr

// #include <stdio.h>
// #include <tesseract/capi.h>
import "C"
import "sync"

type Rectangle struct {
	left   int
	top    int
	width  int
	height int
}

type TesseractAPI struct {
	api      *C.struct_TessBaseAPI
	disposed bool
	mutex    sync.Mutex
}
