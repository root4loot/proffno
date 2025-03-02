package proffno

import (
	"testing"
)

func TestGetSubsidiaries(t *testing.T) {
	orgName := "DnB Bank ASA"
	greaterThanPercentage := 50.0

	results, err := GetSubsidiaries(orgName, greaterThanPercentage)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedSubsidiary := "DnB Asset Management AS"
	found := false
	for _, result := range results {
		if result.Name == expectedSubsidiary {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("Expected to find %s in subsidiaries, but it was not found", expectedSubsidiary)
	}
}
