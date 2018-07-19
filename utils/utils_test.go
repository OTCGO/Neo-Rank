package utils

import (
	"testing"
)

func TestAddressToScripthash(t *testing.T) {

	expected := "0xa7274594ce215208c8e309e8f2fe05d4a9ae412b"
	actual := BigOrLittle(AddressToScripthash("AKibPRzkoZpHnPkF6qvuW2Q4hG9gKBwGpR"))

	if actual != expected {
		t.Errorf("Expected the  to be %s but instead got %s!", expected, actual)
	}
}
