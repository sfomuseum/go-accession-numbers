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
go build -mod vendor -o bin/flatten-definition cmd/flatten-definition/main.go
```

### flatten-definition

```
> ./bin/flatten-definition -h
Usage of ./bin/flatten-definition:
  -constvar constvar
    	Encode the output as a valid gocloud.dev/runtimevar constvar string.
  -path string
    	The path your sfomuseum/accession-numbers definition file.
```

For example:

```
$> ./bin/flatten-definition \
	-path /usr/local/data/accession-numbers/data/sfomuseum.org.json \
	-constvar 

constant://?decoder=string&val=%7B%22organization_name%22%3A%22SFO+Museum%22%2C%22organization_url%22%3A%22https%3A%2F%2Fsfomuseum.org%2F%22%2C%22object_url%22%3A%22https%3A%2F%2Fcollection.sfomuseum.org%2Fobjects%2F%7Baccession_number%7D%22%2C%22iiif_manifest%22%3A%22https%3A%2F%2Fcollection.sfomuseum.org%2Fobjects%2F%7Baccession_number%7D%2Fmanifest%22%2C%22oembed_profile%22%3A%22https%3A%2F%2Fcollection.sfomuseum.org%2Foembed%2F%3Furl%3Dhttps%3A%2F%2Fcollection.sfomuseum.org%2Fobjects%2F%7Baccession_number%7D%5Cu0026format%3Djson%22%2C%22whosonfirst_id%22%3A102527513%2C%22patterns%22%3A%5B%7B%22label%22%3A%22common%22%2C%22pattern%22%3A%22%28%28%3F%3AL%7CR%29%3F%28%3F%3A%5C%5Cd%2B%29%5C%5C.%28%3F%3A%5C%5Cd%2B%29%5C%5C.%28%3F%3A%5C%5Cd%2B%29%28%3F%3A%5C%5C.%28%3F%3A%5C%5Cd%2B%29%29%3F%28%3F%3A%28%3F%3A%5C%5Cs%3F%5Bsa-z%5D%29%2B%29%3F%29%22%2C%22tests%22%3A%7B%221994.18.175%22%3A%5B%221994.18.175%22%5D%2C%221994.18.199a%22%3A%5B%221994.18.199a%22%5D%2C%222000.058.1185+a+c%22%3A%5B%222000.058.1185+a+c%22%5D%2C%222001.106.041+a%22%3A%5B%222001.106.041+a%22%5D%2C%222002.135.017.042%22%3A%5B%222002.135.017.042%22%5D%2C%222014.120.001%22%3A%5B%222014.120.001%22%5D%2C%22L2021.0501.033+a%22%3A%5B%22L2021.0501.033+a%22%5D%2C%22R2021.0501.030%22%3A%5B%22R2021.0501.030%22%5D%2C%22This+is+an+object%5C%5CnGift+of+Important+Donor%5C%5Cn1994.18.175%5C%5Cn%5C%5CnThis+is+another+object%5C%5CnAnonymouts+Gift%5C%5Cn1994.18.165+1994.18.199a%5C%5Cn2000.058.1185+a+c%5C%5CnOil+on+canvas%22%3A%5B%221994.18.175%22%2C%221994.18.165%22%2C%221994.18.199a%22%2C%222000.058.1185+a+c%22%5D%7D%7D%5D%7D
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
| TWILIO_DEFINITION_URI | string | A valid `gocloud.dev/runtimevar` URI. In a Lambda context this would mean a `constvar://` URI that can be generated using the `flatten-definition` tool described above. |

#### API Gateway

## See also

* https://github.com/sfomuseum/accession-numbers