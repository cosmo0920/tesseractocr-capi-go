package tesseractocr

// #include <stdio.h>
// #include <tesseract/capi.h>
import "C"

type Rectangle struct {
	left int
	top int
	width int
	height int
}

func (rect *Rectangle) SetRectangle(left int, top int, width int, height int) {
	rect.left = left
	rect.top = top
	rect.width = width
	rect.height = height
}

type TesseractAPI struct {
	api *C.struct_TessBaseAPI
}