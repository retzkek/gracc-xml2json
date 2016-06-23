package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"

	"github.com/opensciencegrid/gracc-collector/gracc"
)

func main() {
	flag.Usage = usage
	flag.Parse()

	xml2json(flag.Args())
}

func xml2json(args []string) {
	var fi, fo *os.File
	var err error
	switch len(args) {
	case 0:
		fi = os.Stdin
		fo = os.Stdout
	case 1:
		if fi, err = os.Open(args[0]); err != nil {
			panic(err)
		}
		defer fi.Close()
		fo = os.Stdout
	case 2:
		if fi, err = os.Open(args[0]); err != nil {
			panic(err)
		}
		defer fi.Close()
		if fo, err = os.Create(args[1]); err != nil {
			panic(err)
		}
		defer fo.Close()
	}

	var v gracc.JobUsageRecord
	var buf bytes.Buffer
	if _, err := buf.ReadFrom(fi); err != nil {
		panic(err)
	}
	if err := v.ParseXML(buf.Bytes()); err != nil {
		panic(err)
	}
	if j, err := v.ToJSON("    "); err != nil {
		panic(err)
	} else {
		if _, err := fo.Write(j); err != nil {
			panic(err)
		}
	}

}

func usage() {
	fmt.Printf(`Usage: %s [INPUT] [OUTPUT]

INPUT	input XML file 
OUTPUT	output JSON file 

If no arguments are provided, input is read from stdin
and output is to stdout.

If one argument is specified, it is assumed to be INPUT, 
with output sent to stdout.
`, os.Args[0])
}
