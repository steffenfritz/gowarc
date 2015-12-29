// writer writes data to warc file

package main

import "os"

// create creates the output warc file
func create(warcfile *string) (*os.File, error) {

	f, err := os.Create(*warcfile)
	defer f.Close()

	return f, err
}

// writeInitInfoBlock writes the first block of a warc file
func writeInitInfoBlock(warcfile *os.File) error {

	_, err := warcfile.WriteString("WARC/1.0\r\n")

	return err
}

// writeResponseBlock writes blocks for each file in sorce dir
func writeResponseBlock(warcfile *os.File, payload []byte) error {

	return nil
}
