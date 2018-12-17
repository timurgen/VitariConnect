package main

import "testing"

func TestHeaderA(t *testing.T) {
	expected := "<Header><ClientId>9999</ClientId><Guid>XXXX-XXXX-XXXX-XXXX-XXXX</Guid></Header>"

	actual := BuildHeader("XXXX-XXXX-XXXX-XXXX-XXXX", 9999)
	if actual != expected {
		t.Errorf("header representation incorrect\r\n got %s\r\nexpected %s", actual, expected)
	}
}
