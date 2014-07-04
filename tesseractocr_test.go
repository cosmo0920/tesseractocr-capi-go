package tesseractocr_test

import (
	ocr "./tesseractocr"
	"testing"
	)

func TestBaseAPINew(t *testing.T) {
	result := ocr.BaseAPINew()
	if result == nil {
		t.Errorf("result = %v cannot nil", result, nil)
	}
}

func TestVersion(t *testing.T) {
	result := ocr.Version()
	if result == "" {
		t.Errorf("result = %v cannot empty string", result, nil)
	}
}

func TestEnv(t *testing.T) {
	result := ocr.Env()
	if result == "" {
		t.Errorf("result = %v cannot empty string", result, nil)
	}
}
