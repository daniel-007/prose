/*
Package util contains internals used across the other prose packages.
*/
package util

import (
	"io/ioutil"
	"path/filepath"
)

// ReadDataFile reads data from a file, panicking on any errors.
func ReadDataFile(path string) []byte {
	p, err := filepath.Abs(path)
	CheckError(err)

	data, ferr := ioutil.ReadFile(p)
	CheckError(ferr)

	return data
}

// CheckError panics if `err` is not `nil`.
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
