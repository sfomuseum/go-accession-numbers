package accessionnumbers

import (
	"testing"
)

func TestDefinition(t *testing.T) {

	// defined in tests.go

	def, err := loadTestDefinition()

	if err != nil {
		t.Fatal(err)
	}

	num := "R2021.0501.030"

	expected_iiif_uri := "https://millsfield.sfomuseum.org/objects/R2021.0501.030/manifest"
	expected_oembed_uri := "https://millsfield.sfomuseum.org/oembed/?url=https://millsfield.sfomuseum.org/objects/R2021.0501.030&format=json"
	expected_object_uri := "https://millsfield.sfomuseum.org/objects/R2021.0501.030"

	iiif_uri, err := def.IIIFManifestURI(num)

	if err != nil {
		t.Fatalf("Failed to derive IIIF manifest URI, %v", err)
	}

	if iiif_uri != expected_iiif_uri {
		t.Fatalf("Invalid IIIF URI. Expected '%s' but got '%s'", expected_iiif_uri, iiif_uri)
	}

	oembed_uri, err := def.OEmbedProfileURI(num)

	if err != nil {
		t.Fatalf("Failed to derive OEmbed profile URI, %v", err)
	}

	if oembed_uri != expected_oembed_uri {
		t.Fatalf("Invalid OEmbed profile URI. Expected '%s' but got '%s'", expected_oembed_uri, oembed_uri)
	}

	object_uri, err := def.ObjectURI(num)

	if err != nil {
		t.Fatalf("Failed to derive object URI, %v", err)
	}

	if object_uri != expected_object_uri {
		t.Fatalf("Invalid OBJECT URI. Expected '%s' but got '%s'", expected_object_uri, object_uri)
	}

}
