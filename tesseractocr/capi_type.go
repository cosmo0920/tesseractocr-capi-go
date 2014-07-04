package tesseractocr

// #include <stdio.h>
// #include <tesseract/capi.h>
import "C"

type Rectangle struct {
	left C.int
	top C.int
	width C.int
	height C.int
}

type TesseractAPI struct {
	api *C.struct_TessBaseAPI
}
