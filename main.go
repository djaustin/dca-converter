package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/jonas747/dca"
)

func main() {
	input := flag.String("in", "", "relative path to file to be converted")
	output := flag.String("out", "", "relative path of the output DCA file")
	flag.Parse()
	if *input == "" || *output == "" {
		flag.Usage()
		return
	}

	opts := dca.StdEncodeOptions
	// Do not add metadata - DCA1 not required
	opts.RawOutput = true
	encoding, err := dca.EncodeFile(*input, opts)
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer encoding.Cleanup()

	outputFile, err := os.Create(*output)
	if err != nil {
		fmt.Println("cannot create file", err.Error())
		return
	}

	_, err = io.Copy(outputFile, encoding)
	if err != nil {
		fmt.Println(err.Error())
	}

}
