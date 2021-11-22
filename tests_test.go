package accessionnumbers

import (
	"testing"
)

func TestLoadTestDefinitions(t *testing.T) {

	_, err := loadTestDefinition()

	if err != nil {
		t.Fatalf("Failed to load test definition, %v", err)
	}
}
