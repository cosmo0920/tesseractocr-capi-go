// +build linux

package tesseractocr

// #cgo LDFLAGS: -ltesseract -llept
import "C"

/*
 NOTE: specify LDFLAGS by hand for workaround.
 BUG:1329530 https://bugs.launchpad.net/ubuntu/+source/tesseract/+bug/1329530
*/
