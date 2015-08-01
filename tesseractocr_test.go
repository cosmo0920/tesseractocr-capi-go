package tesseractocr_test

import (
	ocr "."
	"testing"
)

func setupTesseractAPI(t *testing.T) *ocr.TesseractAPI {
	lang := "eng"
	env := ocr.Env()
	api, err := ocr.BaseAPIInit3(env, lang)
	if err != nil {
		t.Errorf("err = %v initialize failure", err)
	}
	return api
}

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

func TestBaseAPISetVariable(t *testing.T) {
	api := setupTesseractAPI(t)
	value := "0123456789"
	api.BaseAPISetVariable("tessedit_char_whitelist", value)
	result, err := api.BaseAPIGetStringVariable("tessedit_char_whitelist")
	if err != nil {
		t.Errorf("result == %v assertion failed", result)
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
	api := setupTesseractAPI(t)
	api.BaseAPISetSegMode(ocr.PSM_SINGLE_WORD)
	mode := api.BaseAPIGetSegMode()
	if mode != ocr.PSM_SINGLE_WORD {
		t.Errorf("mode = %v cannot default (PSM_SINGLE_BLOCK)", mode)
	}
}
