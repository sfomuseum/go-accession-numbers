// twilio-handler provides an HTTP server to listen for and respond to Twilio-style webhook URLs. This server can be run locally or as an AWS Lambda function.
package main

/*

$> go run -mod vendor cmd/twilio-handler/main.go -definition-uri 'file:///usr/local/sfomuseum/accession-numbers/data/sfomuseum.org.json'
2021/12/20 12:01:11 Listening on http://localhost:8080

$> curl -X POST -H 'Content-type: application/x-www-form-urlencoded' -d 'Body=Hello world 1994.18.175' http://localhost:8080
The following objects were found in that text:
https://collection.sfomuseum.org/objects/1994.18.175

*/

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/aaronland/go-http-server"
	"github.com/sfomuseum/go-accession-numbers"
	"github.com/sfomuseum/go-flags/flagset"
	"github.com/sfomuseum/runtimevar"
	_ "gocloud.dev/runtimevar/constantvar"
	_ "gocloud.dev/runtimevar/filevar"
)

/*
Twilio makes HTTP requests to your application, just like a regular web browser, in the format application/x-www-form-urlencoded. By including parameters and values in its requests, Twilio sends data to your application that you can act upon before responding.

https://www.twilio.com/docs/messaging/guides/webhook-request
*/

func AccessionNumbersHandler(def *accessionnumbers.Definition) http.Handler {

	fn := func(rsp http.ResponseWriter, req *http.Request) {

		req.ParseForm()

		f := req.Form
		text := f.Get("Body")

		if text == "" {
			return
		}

		var body string

		matches, err := accessionnumbers.ExtractFromText(text, def)
		uris := make([]string, 0)

		if err != nil {
			log.Printf("Failed to extract text from '%s', %v", body, err)

		} else {

			for _, m := range matches {

				acc_num := m.AccessionNumber
				object_uri, err := def.ObjectURI(acc_num)

				if err != nil {
					log.Printf("Failed to derive object URI for %s, %v", acc_num, err)
					continue
				}

				uris = append(uris, object_uri)
			}
		}

		if len(uris) == 0 {
			body = fmt.Sprintf("There was a problem parsing that text.")
		} else {
			body = fmt.Sprintf("The following objects were found in that text:\n%s", strings.Join(uris, "\n"))
		}

		rsp.Write([]byte(body))
		return
	}

	return http.HandlerFunc(fn)
}

func main() {

	fs := flagset.NewFlagSet("twilio")

	server_uri := fs.String("server-uri", "http://localhost:8080", "A valid aaronland/go-http-server URI.")
	definition_uri := fs.String("definition-uri", "", "A valid gocloud.dev/runtimevar URI. Supported URI schemes are: constant://, file://")

	fs.Usage = func() {
		fmt.Fprintf(os.Stderr, "twilio-handler provides an HTTP server to listen for and respond to Twilio-style SMS webhook URLs. This server can be run locally or as an AWS Lambda function.\n")
		fmt.Fprintf(os.Stderr, "Usage:\n\t %s [options\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Valid options are:\n")
		fs.PrintDefaults()
	}

	flagset.Parse(fs)

	err := flagset.SetFlagsFromEnvVars(fs, "TWILIO")

	if err != nil {
		log.Fatalf("Failed to set flags from environment variables, %v", err)
	}

	ctx := context.Background()

	definition_str, err := runtimevar.StringVar(ctx, *definition_uri)

	if err != nil {
		log.Fatalf("Failed to parse '%s', %v", *definition_uri, err)
	}

	definition_r := strings.NewReader(definition_str)

	var def *accessionnumbers.Definition

	dec := json.NewDecoder(definition_r)
	err = dec.Decode(&def)

	if err != nil {
		log.Fatalf("Failed to decode '%s', %v", *definition_uri, err)
	}

	handler := AccessionNumbersHandler(def)

	mux := http.NewServeMux()
	mux.Handle("/", handler)

	s, err := server.NewServer(ctx, *server_uri)

	if err != nil {
		log.Fatalf("Failed to create new server, %v", err)
	}

	log.Printf("Listening on %s\n", s.Address())

	err = s.ListenAndServe(ctx, mux)

	if err != nil {
		log.Fatalf("Failed to serve requests, %v", err)
	}
}
