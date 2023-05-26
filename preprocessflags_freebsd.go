//go:build cgo
// +build cgo

package gosseract

// #cgo LDFLAGS: -L/usr/local/lib -llept -ltesseract -fopenmp
import "C"
