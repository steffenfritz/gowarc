// reader reads content from a directory

package main

import (
	"os"
	"path/filepath"
)

func test(path string, info os.FileInfo, err error) error {

	return nil
}

func readerWriter(sourceDir string, uri string) {

	filepath.Walk(sourceDir, test)
}
