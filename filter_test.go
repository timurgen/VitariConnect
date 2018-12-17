package main

import "testing"

func TestFilterA(t *testing.T){
	row := CreateFilterRow("OrgUnit1", "GreaterThan", "10000", "")
	expected := "<Filters><OrgUnit1 Compare=\"GreaterThan\" Value1=\"10000\" Operator=\"\"/></Filters>"
	filter := Filter{}

	AddRowToFilter(&filter, row)

	filterStr := RenderFilter(filter)

	if filterStr != expected {
		t.Errorf("rendered filter representation incorrect\r\n got %s\r\nexpected %s", filterStr, expected)
	}
}

func TestFilterB(t *testing.T){
	expected := "<Filters><Year Compare=\"EqualTo\" Value1=\"2017\" Operator=\"\"/><Period Compare=\"EqualTo\" Value1=\"11\" Operator=\"And\"/></Filters>"

	rowA := CreateFilterRow("Year", "EqualTo", "2017", "")
	rowB := CreateFilterRow("Period", "EqualTo", "11", "And")

	filter := Filter{}

	AddRowToFilter(&filter, rowA)
	AddRowToFilter(&filter, rowB)

	filterStr := RenderFilter(filter)

	if filterStr != expected {
		t.Errorf("rendered filter representation incorrect\r\n got %s\r\nexpected %s", filterStr, expected)
	}
}
