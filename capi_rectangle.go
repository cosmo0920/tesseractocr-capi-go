package tesseractocr

import "C"

func RectangleNew(left int, top int, width int, height int) (rect *Rectangle) {
	rect.left = left
	rect.top = top
	rect.width = width
	rect.height = height

	rec := &Rectangle{rect.left, rect.top, rect.width, rect.height}
	return rec
}
