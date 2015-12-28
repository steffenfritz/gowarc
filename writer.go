// writer writes data to warc file

package main

import "os"

func writer(warcfile *string) (*os.File, error) {

	f, err := os.Create(*warcfile)
	defer f.Close()

	return f, err
}
