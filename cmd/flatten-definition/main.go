// flatten-definition is a tool to "flatten" a valid sfomuseum/accession-numbers definition file in to a string that can be copy-pasted
// in to an (AWS) Lambda environment variable field.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/url"
	"os"
	"path/filepath"

	"github.com/sfomuseum/go-accession-numbers"
)

func main() {

	path := flag.String("path", "", "The path your sfomuseum/accession-numbers definition file.")
	constvar := flag.Bool("constvar", false, "Encode the output as a valid gocloud.dev/runtimevar `constvar` string.")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "\"Flatten\" an accession number defintion file by removing all newlines and non-relevant whitespace, optionally encoding the result as a gocloud.dev/runtimevar constvar string.\n")
		fmt.Fprintf(os.Stderr, "Usage:\n\t %s [options\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Valid options are:\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	abs_path, err := filepath.Abs(*path)

	if err != nil {
		log.Fatalf("Failed to derive absolute path for '%s', %v", *path, err)
	}

	_, err = os.Stat(abs_path)

	if err != nil {
		log.Fatalf("Failed to stat '%s', %v", abs_path, err)
	}

	fh, err := os.Open(abs_path)

	if err != nil {
		log.Fatalf("Failed to open '%s' for reading, %v", abs_path, err)
	}

	defer fh.Close()

	body, err := io.ReadAll(fh)

	if err != nil {
		log.Fatal(err)
	}

	var def *accessionnumbers.Definition

	err = json.Unmarshal(body, &def)

	if err != nil {
		log.Fatalf("Failed to parse '%s', %v", abs_path, err)
	}

	body, err = json.Marshal(def)

	if err != nil {
		log.Fatalf("Failed to encode '%s', %v", abs_path, err)
	}

	str_body := string(body)

	if *constvar {

		q := url.Values{}
		q.Set("decoder", "string")
		q.Set("val", str_body)

		str_body = fmt.Sprintf("constant://?%s", q.Encode())
	}

	fmt.Println(str_body)
}
