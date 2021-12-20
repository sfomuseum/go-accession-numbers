# go-accession-numbers

Go package providing methods for identifying and extracting accession numbers from arbitrary bodies of text.

## Documentation

[![Go Reference](https://pkg.go.dev/badge/github.com/sfomuseum/go-accession-numbers.svg)](https://pkg.go.dev/github.com/sfomuseum/go-accession-numbers)

## Example

_Error handling omitted for the sake of brevity._

### Basic

```
package main

import (
	"fmt"
	"github.com/sfomuseum/go-accession-numbers"
)

func main() {

	re_pat := `((?:L|R)?(?:\d+)\.(?:\d+)\.(?:\d+)(?:\.(?:\d+))?(?:(?:\s?[sa-z])+)?)`
     
	texts := []string{
     		"2000.058.1185 a c",
		"This is an object\nGift of Important Donor\n1994.18.175\n\nThis is another object\nAnonymouts Gift\n1994.18.165 1994.18.199a\n2000.058.1185 a c\nOil on canvas",
     	}

	for _, t := range texts {
		     
     		m, _ := accessionnumbers.FindMatches(text, re)

		for _, m := range matches {
			fmt.Printf("%s\n", m)
		}
     	}
```

This would yield:

```
2000.058.1185 a c
1994.18.175
1994.18.165
1994.18.199a
2000.058.1185 a c
```

### Using "defintion" files

"Definition" files are provided by the [sfomuseum/accession-numbers](https://github.com/sfomuseum/accession-numbers) package.

```
package main

import (
	"encoding/json"
	"fmt"		
	"github.com/sfomuseum/go-accession-numbers"
	"os"
)

func main() {

	var def *Definition
	
	r, _ := os.Open("fixtures/sfomuseum.json")

	dec := json.NewDecoder(r)
	dec.Decode(&def)

	re_pat := `((?:L|R)?(?:\d+)\.(?:\d+)\.(?:\d+)(?:\.(?:\d+))?(?:(?:\s?[sa-z])+)?)`
     
	texts := []string{
     		"2000.058.1185 a c",
		"This is an object\nGift of Important Donor\n1994.18.175\n\nThis is another object\nAnonymouts Gift\n1994.18.165 1994.18.199a\n2000.058.1185 a c\nOil on canvas",
     	}

	for _, t := range texts {

		matches, _ := accessionnumbers.ExtractFromText(t, def)
		
		for _, m := range matchess {
			fmt.Printf("%s (%s)\n", m.AccessionNumber, m.OrganizationURL)
		}
     	}
```

This would yield:

```
2000.058.1185 a c (https://sfomuseum.org/)
1994.18.175 (https://sfomuseum.org/)
1994.18.165 (https://sfomuseum.org/)
1994.18.199a (https://sfomuseum.org/)
2000.058.1185 a c (https://sfomuseum.org/)
```

## Tools

```
$> make cli
go build -mod vendor -o bin/twilio-handler cmd/twilio-handler/main.go
```

### twilio-handler

```
$> ./bin/twilio-handler -h
  -definition-uri string
    	A valid gocloud.dev/runtimevar URI. (default "accessionnumbers/sfomuseum.org.json")
  -server-uri string
    	A valid aaronland/go-http-server URI. (default "http://localhost:8080")
```

For example, running the application locally:

```
$> ./bin/twilio-handler -definition-uri 'file:///usr/local/sfomuseum/accession-numbers/data/sfomuseum.org.json'
2021/12/20 12:01:11 Listening on http://localhost:8080

$> curl -X POST -H 'Content-type: application/x-www-form-urlencoded' -d 'Body=Hello world 1994.18.175' http://localhost:8080
The following objects were found in that text:
https://collection.sfomuseum.org/objects/1994.18.175
```

### AWS

#### Lambda

```
$> make lambda-twilio-handler
if test -f main; then rm -f main; fi
if test -f twilio-handler.zip; then rm -f twilio-handler.zip; fi
GOOS=linux go build -mod vendor -o main cmd/twilio-handler/main.go
zip twilio-handler.zip main
  adding: main (deflated 58%)
rm -f main

$> ls -la twilio-handler.zip 
-rw-r--r--  1 asc  staff  11479660 Dec 20 12:08 twilio-handler.zip
```

##### Environment variables

| Name | Value | Notes
| --- | --- | --- |
| TWILIO_SERVER_URI | `lambda://` | |
| TWILIO_DEFINITION_URI | string | | ... |

#### API Gateway

## See also

* https://github.com/sfomuseum/accession-numbers