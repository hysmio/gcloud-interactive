package parser

import (
	"testing"

	"github.com/hysmio/gcloud-interactive/parser"
)

func TestParseString(t *testing.T) {
	testString := "PROJECT_ID                     NAME             PROJECT_NUMBER\n"
	testString += "automated-cloud-1554084223154  Automated Cloud  576202887291\n"
	testString += "some junk project id           Stupid name\n"
	testString += "                               No ID            371982723874\n"

	values := parser.ParseTable(testString)

	template := "%s expected: \"%s\" | actual: \"%s\""

	// Test Row 0
	row := 0

	expected := "automated-cloud-1554084223154"
	header := "PROJECT_ID"
	if values[row][header] != expected {
		t.Errorf(template, header, expected, values[row][header])
	}

	expected = "Automated Cloud"
	header = "NAME"
	if values[row][header] != expected {
		t.Errorf(template, header, expected, values[row][header])
	}

	expected = "576202887291"
	header = "PROJECT_NUMBER"
	if values[row][header] != expected {
		t.Errorf(template, header, expected, values[row][header])
	}

	// Test Row 1
	row = 1

	expected = "some junk project id"
	header = "PROJECT_ID"
	if values[row][header] != expected {
		t.Errorf(template, header, expected, values[row][header])
	}

	expected = "Stupid name"
	header = "NAME"
	if values[row][header] != expected {
		t.Errorf(template, header, expected, values[row][header])
	}

	expected = ""
	header = "PROJECT_NUMBER"
	if values[row][header] != expected {
		t.Errorf(template, header, expected, values[row][header])
	}

	// Test Row 2
	row = 2

	expected = ""
	header = "PROJECT_ID"
	if values[row][header] != expected {
		t.Errorf(template, header, expected, values[row][header])
	}

	expected = "No ID"
	header = "NAME"
	if values[row][header] != expected {
		t.Errorf(template, header, expected, values[row][header])
	}

	expected = "371982723874"
	header = "PROJECT_NUMBER"
	if values[row][header] != expected {
		t.Errorf(template, header, expected, values[row][header])
	}
}
