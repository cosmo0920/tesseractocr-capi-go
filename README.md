TesseractOCR C API Go wrapper
===

[![Build Status](https://travis-ci.org/cosmo0920/tesseractocr-capi-go.svg?branch=master)](https://travis-ci.org/cosmo0920/tesseractocr-capi-go) [![GoDoc](https://godoc.org/github.com/cosmo0920/tesseractocr-capi-go/tesseractocr?status.png)](https://godoc.org/github.com/cosmo0920/tesseractocr-capi-go/tesseractocr)

## instalation

Before you use this library, please install `tesseract-ocr`.

### OSX

For OSX 10.9 Marvericks with homebrew:

```bash
$ brew install tesseract-ocr --all-languages --with-tiff
```

### Ubuntu Linux

For Ubuntu Linux 14.04 LTS (Trusty Tahr):

```bash
$ sudo apt-get install libtesserac-dev libleptonica-dev
$ sudo apt-get install tesseract-ocr-eng tesseract-ocr-jpn
```

### FreeBSD (pkgng)

FreeBSD 10 release with pkgng:

```bash
$ export CC=clang
$ sudo pkg install -y tesseract pkgconf
$ sudo pkg install -y tesseract-data
```

## import

```bash
$ go get "github.com/cosmo0920/tesseractocr-capi-go/tesseractocr"
```

and then,

```go
import "github.com/cosmo0920/tesseractocr-capi-go/tesseractocr"
```

## LICENSE

[Apache License 2.0](LICENSE)

## Support platform

* OSX 10.9 Marvericks with homebrew
* Ubuntu Linux 14.04 LTS (Trusty Tahr)
* FreeBSD 10 release with pkgng
