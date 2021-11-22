package accessionnumbers

import (
	"fmt"
	"testing"
)

func TestExtractFromText(t *testing.T) {

	// defined in tests.go

	def, err := loadTestDefinition()

	if err != nil {
		t.Fatal(err)
	}

	err = testDefinition(def)

	if err != nil {
		t.Fatal(err)
	}
}

func testDefinition(def *Definition) error {

	for _, p := range def.Patterns {

		err := testPattern(p)

		if err != nil {
			return err
		}
	}

	return nil
}

func testPattern(p *Pattern) error {

	for str, expected := range p.Tests {

		fmt.Printf("Find matches for '%s' using '%s'\n", str, p.Pattern)

		m, err := FindMatches(str, p.Pattern)

		if err != nil {
			return fmt.Errorf("Failed to find accession numbers for '%s' using '%s', %v", str, p.Label, err)
		}

		if len(m) != len(expected) {
			return fmt.Errorf("Invalid count for '%s' using '%s', expected %d results but got %d", str, p.Label, len(expected), len(m))
		}
	}

	return nil
}
