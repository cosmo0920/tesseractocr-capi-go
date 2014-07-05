TesseractOCR C API Go wrapper
===

[![Build Status](https://travis-ci.org/cosmo0920/tesseractocr-capi-go.svg?branch=master)](https://travis-ci.org/cosmo0920/tesseractocr-capi-go) [![GoDoc](https://godoc.org/github.com/cosmo0920/tesseractocr-capi-go/tesseractocr?status.png)](https://godoc.org/github.com/cosmo0920/tesseractocr-capi-go/tesseractocr)

## instalation

Before you use this library, please install `tesseract-ocr`.

### OSX

```bash
$ brew install tesseract-ocr --all-languages --with-tiff
```

### Ubuntu Linux

```bash
$ sudo aptget-install tesseractocr-dev libleptonica-dev
```

### FreeBSD

```bash
$ export CC=clang
$ sudo pkg install -y tesseract pkgconf
```

## import

```bash
$ go get "github.com/cosmo0920/tesseractocr-capi-go/tesseractocr"
```

or

```go
import "github.com/cosmo0920/tesseractocr-capi-go/tesseractocr"
```

## LICENSE

[Apache License 2.0](LICENSE)
