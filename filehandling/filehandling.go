package filehandling

import (
	"os"
)

//StatFile is a simple check, if a given file (string path to file) is existent and is a regular readable file
func StatFile(file string) bool {
	fileObj, fileErr := os.Stat(file)
	if fileErr != nil {
		return false
	}
	if !fileObj.Mode().IsRegular() {
		return false
	}
	return true
}
