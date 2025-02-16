package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	type test struct {
		name     string
		input    http.Header
		expected string
	}

	testHeader := make(http.Header)
	testHeader.Add("Authorization", "ApiKey asdf")

	tests := []test{
		{name: "Check for ApiKey in Header", input: testHeader, expected: "asdf"},
	}

	for _, tc := range tests {
		actual, err := GetAPIKey(testHeader)
		t.Logf("Running test %s...", tc.name)
		if err != nil {
			t.Logf("Test Failed: %v", err)
		}
		if !reflect.DeepEqual(tc.expected, actual) {
			t.Fatalf("expected: %v, got: %v", tc.expected, actual)
		}
	}
}
