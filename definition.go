package accessionnumbers

import (
	"fmt"
	"net/url"
)

// type Definition provides a struct containing accession number patterns and URIs for an organization.
type Definition struct {
	// The name of the organization associated with this definition.
	OrganizationName string `json:"organization_name"`
	// The URL of the organization associated with this definition.
	OrganizationURL string `json:"organization_url"`
	// A valid URI template (RFC 6570) used to generate the URL for an object given its accession number.
	ObjectURLTemplate string `json:"object_url,omitempty"`
	// A valid URI template (RFC 6570) used to generate the IIIF manifest URL for an object given its accession number.
	IIIFManifestTemplate string `json:"iiif_manifest,omitempty"`
	// A valid URI template (RFC 6570) used to generate an OEmbed profile URL for an object given its accession number.
	OEmbedProfileTemplate string `json:"oembed_profile,omitempty"`
	// A valid Who's On First ID representing the organization.
	WhosOnFirstId int64 `json:"whosonfirst_id,omitempty"`
	// The set of patterns used to identify and extract accession numbers associated with an organization.
	Patterns []*Pattern `json:"patterns"`
}

func (def *Definition) IIIFManifestURL(accession_number string) (*url.URL, error) {

	if def.IIIFManifestTemplate == "" {
		return nil, fmt.Errorf("IIIFManifestURLTemplate is undefined")
	}

	return nil, fmt.Errorf("Not implemented")
}

func (def *Definition) OEmbedProfileURL(accession_number string) (*url.URL, error) {

	if def.OEmbedProfileTemplate == "" {
		return nil, fmt.Errorf("OEmbedProfileURLTemplate is undefined")
	}

	return nil, fmt.Errorf("Not implemented")
}

func (def *Definition) ObjectURL(accession_number string) (*url.URL, error) {

	if def.ObjectURLTemplate == "" {
		return nil, fmt.Errorf("ObjectURLTemplate is undefined")
	}

	return nil, fmt.Errorf("Not implemented")
}
