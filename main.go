package main

import (
	"fmt"
	"flag"
	"io/ioutil"
	"text/scanner"
)

var inFile = flag.String( "in_file", "", "Specify the file to read from")
var outFile = flag.String ( "out_file", "", "Specify the file to write to"  )

func main() {
	if *inFile == "" || *outFile == "" {
		fmt.Println("Error, please specify the files")
	}
	data, err := ioutil.ReadFile( *inFile )
	if err != nil {
		var s scanner.Scanner
		s.Init( data )
		for token := s.Scan(); token != scanner.EOF; {

		}
	}
}

