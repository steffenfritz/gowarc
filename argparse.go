// argparse parses arguments
// TODO: Error handling

package main

import (
	"errors"
	"flag"
)

func argparse() (*string, *string, *string, error) {

	var source = flag.String("i", "", "name of source directory")
	var target = flag.String("o", "", "name of target warc file")
	var uri = flag.String("u", "", "uri used as target-uri")

	flag.Parse()

	if flag.NFlag() != 3 {
		err := errors.New("USAGE: gowarc -i=$SOURCEDIR -o=$OUTPUTWARC -u=$Target-URI")
		return source, target, uri, err
	}
	return source, target, uri, nil
}
