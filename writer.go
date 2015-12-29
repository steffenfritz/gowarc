// writer writes data to warc file

package main

import "os"

func create(warcfile *string) (*os.File, error) {

	f, err := os.Create(*warcfile)
	defer f.Close()

	return f, err
}

func writeInitInfoBlock(warcfile *os.File) error {

	return nil
}

func writeResponseBlock(warcfile *os.File, payload []byte) error {

	return nil
}
