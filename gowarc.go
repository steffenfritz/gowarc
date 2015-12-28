/*
Copyright (C) 2015  Steffen Fritz

This program is free software; you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation; either version 2 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License along
with this program; if not, write to the Free Software Foundation, Inc.,
51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA.

eMail: steffen AT fritz DOT wtf
*/

// gowarc creates a warc file from arbitrary files

package main

import (
	"log"
)

func errorcheck(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {

	source, target, uri, err := argparse()
	errorcheck(err)

	f, err = writer(target)
	errorcheck(err)
}
