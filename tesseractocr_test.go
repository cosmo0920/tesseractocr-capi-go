package tesseractocr_test

import (
	ocr "./tesseractocr"
	"testing"
)

func TestBaseAPIInit2(t *testing.T) {
	lang := "eng"
	env := ocr.Env()
	oem := ocr.OEM_DEFAULT
	_, err := ocr.BaseAPIInit2(env, lang, oem)
	if err != nil {
		t.Errorf("err = %v initialize failure", err)
	}
}

func TestBaseAPIInit3(t *testing.T) {
	lang := "eng"
	env := ocr.Env()
	_, err := ocr.BaseAPIInit3(env, lang)
	if err != nil {
		t.Errorf("err = %v initialize failure", err)
	}
}

func TestVersion(t *testing.T) {
	result := ocr.Version()
	if result == "" {
		t.Errorf("result = %v cannot empty string", result)
	}
}

func TestEnv(t *testing.T) {
	result := ocr.Env()
	if result == "" {
		t.Errorf("result = %v cannot empty string", result)
	}
}

func TestSegMode(t *testing.T) {
		lang := "eng"
	env := ocr.Env()
	api, err := ocr.BaseAPIInit3(env, lang)
	if err != nil {
		t.Errorf("err = %v initialize failure", err)
	}
	api.BaseAPISetSegMode(ocr.PSM_SINGLE_WORD)
	mode := api.BaseAPIGetSegMode()
	if mode != ocr.PSM_SINGLE_WORD {
		t.Errorf("mode = %v cannot default (PSM_SINGLE_BLOCK)", mode)
	}
}
