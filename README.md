TesseractOCR C API Go wrapper
===

[![Build Status](https://travis-ci.org/cosmo0920/tesseractocr-capi-go.svg?branch=master)](https://travis-ci.org/cosmo0920/tesseractocr-capi-go) [![GoDoc](https://godoc.org/github.com/cosmo0920/tesseractocr-capi-go?status.png)](https://godoc.org/github.com/cosmo0920/tesseractocr-capi-go)

## instalation

Before you use this library, please install `tesseract-ocr`.

* tesseract-ocr 3.02 needs for all tests passing
* tesseract-ocr 3.03 also supported but some tests failing
* tesseract-ocr 3.04 is not supported yet.

### OSX

For OSX 10.9 Marvericks with homebrew:

```bash
$ brew install tesseract --all-languages
```

### Ubuntu Linux

For Ubuntu Linux 14.04 LTS (Trusty Tahr):

```bash
$ sudo apt-get install libtesseract-dev libleptonica-dev
$ sudo apt-get install tesseract-ocr-eng tesseract-ocr-jpn
```

### FreeBSD (pkgng)

FreeBSD 10 release with pkgng:

```bash
$ export CC=clang
$ sudo pkg install -y tesseract pkgconf
$ sudo pkg install -y tesseract-data
```

### Windows

Not supported yet.

## import

This library depends on https://github.com/cosmo0920/leptonica-capi-go/.

```bash
$ go get "github.com/cosmo0920/leptonica-capi-go/"
$ go get "github.com/cosmo0920/tesseractocr-capi-go/"
```

and then,

```go
import "github.com/cosmo0920/tesseractocr-capi-go/"
```

## LICENSE

[Apache License 2.0](LICENSE)

## Support platform

* OSX 10.9 Marvericks with homebrew
* Ubuntu Linux 14.04 LTS (Trusty Tahr)
* FreeBSD 10 release with pkgng
