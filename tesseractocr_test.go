package tesseractocr_test

import (
	ocr "./tesseractocr"
	"testing"
)

func TestBaseAPINew(t *testing.T) {
	result := ocr.BaseAPINew()
	if result == nil {
		t.Errorf("result = %v cannot nil", result)
	}
}

func TestBaseAPIInit2(t *testing.T) {
	lang := "eng"
	env := ocr.Env()
	api := ocr.BaseAPINew()
	oem := ocr.OEM_CUBE_ONLY
	_, err := api.BaseAPIInit2(env, lang, oem)
	if err != nil {
		t.Errorf("err = %v initialize failure", err)
	}
}

func TestBaseAPIInit3(t *testing.T) {
	lang := "eng"
	env := ocr.Env()
	api := ocr.BaseAPINew()
	_, err := api.BaseAPIInit3(env, lang)
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
