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

func (rect *Rectangle) SetRectangle(left C.int, top C.int, width C.int, height C.int) {
	rect.left = left
	rect.top = top
	rect.width = width
	rect.height = height
}

type TesseractAPI struct {
	api *C.struct_TessBaseAPI
}