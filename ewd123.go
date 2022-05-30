/*
 * ewd123.go
 *
 * A program to represent the first mutex strategy, as described in EWD123.
 *
 * Copyright (c) 2019-2019 HS Emden/Leer
 * All Rights Reserved.
 *
 * version 1.00 - 21 Oct 2019 - GJV - initial version
 *
 * author: Gert Veltink, gert.veltink@hs-emden-leer.de (GJV)
 */

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hsel-pds/EWD123/ewd123a"
	"github.com/hsel-pds/EWD123/ewd123b"
	"github.com/hsel-pds/EWD123/ewd123c"
	"github.com/hsel-pds/EWD123/ewd123d"
	"github.com/hsel-pds/EWD123/ewd123dekker"
)

// main is the entry point for execution.
func main() {
	var quitChannel = make(chan string)                          // channel to send a quit signal
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile) // set specific logging attributes

	log.Printf("*** Start\n")

	var variant = "d"
	if len(os.Args[1:]) > 0 { // do we have arguments?
		variant = os.Args[1]
	}

	switch variant {
	case "a":
		ewd123a.Start()
	case "b":
		ewd123b.Start()
	case "c":
		ewd123c.Start()
	case "d":
		ewd123d.Start()
	case "dekker":
		ewd123dekker.Start()
	default:
		log.Printf("unknown variant: '%s'!", variant)
		fmt.Printf("Usage: %s <variant>, where variant is a,b,c,d or empty for the Dekker solution\n", os.Args[0])
		os.Exit(-1) // signal error code -1
	}

	log.Printf("*** Finish: %s\n", <-quitChannel) // wait for a quit signal
}
