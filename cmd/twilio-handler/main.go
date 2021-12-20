package main

import (
	"context"
	_ "encoding/json"
	"fmt"
	"github.com/aaronland/go-http-server"
	"github.com/sfomuseum/go-accession-numbers"
	"github.com/sfomuseum/go-flags/flagset"
	"log"
	"net/http"
	"strings"
)

/*
Twilio makes HTTP requests to your application, just like a regular web browser, in the format application/x-www-form-urlencoded. By including parameters and values in its requests, Twilio sends data to your application that you can act upon before responding.
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

	fs := flagset.NewFlagSet("accession numbers")

	server_uri := fs.String("server-uri", "http://localhost:8080", "A valid aaronland/go-http-server URI")
	definition_uri := fs.String("definition-uri", "accessionnumbers/sfomuseum.org.json", "...")

	flagset.Parse(fs)

	err := flagset.SetFlagsFromEnvVars(fs, "TWILIO")

	if err != nil {
		log.Fatalf("Failed to set flags from environment variables, %v", err)
	}

	ctx := context.Background()

	var def *accessionnumbers.Definition

	// TBD...

	log.Println(*definition_uri)
	
	/*
	dec := json.NewDecoder(fh)
	err = dec.Decode(&def)

	if err != nil {
		log.Fatalf("Failed to decode '%s', %v", *definition_uri, err)
	}
	*/
	
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
