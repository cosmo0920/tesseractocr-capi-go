package tesseractocr

// #include <stdio.h>
// #include <stdlib.h>
// #include <tesseract/capi.h>
import "C"

const (
	success = iota
	failure
	uninitialized
	abort
	unknown
)

type TessOcrEngineMode C.TessOcrEngineMode

const (
	OEM_TESSERACT_ONLY TessOcrEngineMode = iota
	OEM_CUBE_ONLY
	OEM_TESSERACT_CUBE_COMBINED
	OEM_DEFAULT
)

type TessPageSegMode C.TessPageSegMode

const (
	PSM_OSD_ONLY TessPageSegMode = iota
	PSM_AUTO_OSD
	PSM_AUTO_ONLY
	PSM_AUTO
	PSM_SINGLE_COLUMN
	PSM_SINGLE_BLOCK_VERT_TEXT
	PSM_SINGLE_BLOCK
	PSM_SINGLE_LINE
	PSM_SINGLE_WORD
	PSM_CIRCLE_WORD
	PSM_SINGLE_CHAR
	PSM_SPARSE_TEXT
	PSM_SPARSE_TEXT_OSD
	PSM_COUNT
)

type TessPageIteratorLevel C.TessPageIteratorLevel

const (
	RIL_BLOCK TessPageIteratorLevel = iota
	RIL_PARA
	RIL_TEXTLINE
	RIL_WORD
	RIL_SYMBOL
)
